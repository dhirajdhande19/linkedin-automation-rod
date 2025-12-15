package browser

import (
	"fmt"
	"math/rand"
	"time"

	"linkedin-automation/internal/auth"
	"linkedin-automation/internal/stealth"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

// StartBrowser launches Chrome and tests the login engine on a mock page
func StartBrowser() (*rod.Browser, *rod.Page) {
	rand.Seed(time.Now().UnixNano())

	// Random realistic viewport
	width := rand.Intn(400) + 1200
	height := rand.Intn(300) + 700

	// System Chrome (Windows fallback)
	chromePath := "C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe"

	url := launcher.New().
		Bin(chromePath).
		Headless(false).
		Set("window-size", fmt.Sprintf("%d,%d", width, height)).
		MustLaunch()

	browser := rod.New().ControlURL(url).MustConnect()

	// Load simple page
	page := browser.MustPage("https://example.com").MustWaitLoad()

	// ------------------------------------------------
	// 1️⃣ Inject MOCK LOGIN HTML (NO SCRIPT HERE)
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

	// ------------------------------------------------
	// 2️⃣ Attach JS logic SEPARATELY (CRITICAL)
	// ------------------------------------------------
	page.MustEval(`
() => {
	const btn = document.getElementById("loginBtn");

	btn.addEventListener("click", () => {
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

	stealth.Think()

	// ------------------------------------------------
	// 3️⃣ LOGIN ENGINE TEST
	// ------------------------------------------------
	loginCfg := auth.LoginConfig{
		UsernameSelector: `input[name="username"]`,
		PasswordSelector: `input[name="password"]`,
		SubmitSelector:   `#loginBtn`,
		SuccessCheckJS:   `() => document.body.innerText.includes("Welcome")`,
		FailureCheckJS:   `() => document.body.innerText.includes("Invalid")`,
		Timeout:          15 * time.Second,
	}

	err := auth.PerformLogin(page, loginCfg, "demo", "demo123")
	if err != nil {
		panic(err)
	}

	// Keep browser open briefly so you can SEE success
	stealth.SleepRandom(3000, 5000)

	return browser, page
}
