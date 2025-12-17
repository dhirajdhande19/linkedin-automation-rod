package browser

import (
	"fmt"
	"math/rand"
	"time"

	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/config"
	"linkedin-automation/internal/session"
	"linkedin-automation/internal/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func StartBrowser() (*rod.Browser, *rod.Page) {
	cfg := config.Load()

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

	// Load cookies if present
	if session.CookiesExist(cfg.CookiePath) {
		_ = session.LoadCookies(browser, cfg.CookiePath)
	}

	// ---------------- MOCK LOGIN PAGE ----------------
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

		if err := auth.PerformLogin(page, loginCfg, cfg.Username, cfg.Password); err != nil {
			panic(err)
		}

		_ = session.SaveCookies(browser, cfg.CookiePath)
	}

	stealth.SleepRandom(3000, 5000)
	return browser, page
}
