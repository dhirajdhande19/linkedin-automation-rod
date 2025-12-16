package browser

import (
	"fmt"
	"math/rand"
	"time"

	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/session"
	"linkedin-automation/internal/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// StartBrowser launches Chrome and tests the login engine on a mock page
func StartBrowser() (*rod.Browser, *rod.Page) {
	cookiePath := "session.json"

	rand.Seed(time.Now().UnixNano())

	width := rand.Intn(400) + 1200
	height := rand.Intn(300) + 700

	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"

	url := launcher.New().
		Bin(chromePath).
		Headless(false).
		Set("window-size", fmt.Sprintf("%d,%d", width, height)).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()
	page := browser.MustPage("https://example.com").MustWaitLoad()

	// Load session cookies (NOTE: mock page resets DOM each run)
	if session.CookiesExist(cookiePath) {
		_ = session.LoadCookies(browser, cookiePath)
	}

	// ------------------------------------------------
	// Inject MOCK LOGIN HTML (POC only)
	// ------------------------------------------------
	page.MustEval(`
() => {
	document.body.innerHTML = 
		` + "`" + `
		<h2>Mock Login</h2>
		<input name="username" placeholder="Username" />
		<br/><br/>
		<input name="password" type="password" placeholder="Password" />
		<br/><br/>
		<button id="loginBtn">Login</button>
		<p id="msg"></p>
		` + "`" + `;
}
`)

	// Attach JS logic
	page.MustEval(`
() => {
	document.getElementById("loginBtn").addEventListener("click", () => {
		const u = document.querySelector('input[name="username"]').value;
		const p = document.querySelector('input[name="password"]').value;

		if (u === "demo" && p === "demo123") {
			document.body.style.background = "#d4edda";
			document.getElementById("msg").innerText = "Welcome";
		} else {
			document.body.style.background = "#f8d7da";
			document.getElementById("msg").innerText = "Invalid";
		}
	});
}
`)

	alreadyLoggedIn := page.MustEval(
		`() => document.body.innerText.includes("Welcome")`,
	).Bool()

	if !alreadyLoggedIn {
		loginCfg := auth.LoginConfig{
			UsernameSelector: `input[name="username"]`,
			PasswordSelector: `input[name="password"]`,
			SubmitSelector:   `#loginBtn`,
			SuccessCheckJS:   `() => document.body.innerText.includes("Welcome")`,
			FailureCheckJS:   `() => document.body.innerText.includes("Invalid")`,
			Timeout:          15 * time.Second,
		}

		if err := auth.PerformLogin(page, loginCfg, "demo", "demo123"); err != nil {
			panic(err)
		}

		// Save cookies after successful login (POC)
		_ = session.SaveCookies(browser, cookiePath)
	}

	stealth.SleepRandom(3000, 5000)
	return browser, page
}
