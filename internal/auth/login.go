package auth

import (
	"errors"
	"time"

	"linkedin-automation/internal/stealth"

	"github.com/go-rod/rod"
)

// LoginConfig defines selectors for any login page
type LoginConfig struct {
	UsernameSelector string
	PasswordSelector string
	SubmitSelector   string
	SuccessCheckJS   string
	FailureCheckJS   string
	Timeout          time.Duration
}

// PerformLogin performs a human-like login flow
func PerformLogin(
	page *rod.Page,
	cfg LoginConfig,
	username string,
	password string,
) error {

	stealth.Think()

	// Find username field
	userEl := page.MustElement(cfg.UsernameSelector)
	userEl.MustScrollIntoView()
	userEl.MustClick()
	stealth.MicroPause()
	stealth.TypeHuman(userEl, username)

	stealth.Think()

	// Find password field
	passEl := page.MustElement(cfg.PasswordSelector)
	passEl.MustScrollIntoView()
	passEl.MustClick()
	stealth.MicroPause()
	stealth.TypeHuman(passEl, password)

	stealth.Think()

	// Click submit
	submit := page.MustElement(cfg.SubmitSelector)
	submit.MustScrollIntoView()

	stealth.Think()

	// Move mouse near submit button
	box := submit.MustShape().Box()
	stealth.MoveMouseHuman(
		page,
		box.X+box.Width/2,
		box.Y+box.Height/2,
	)

	stealth.MicroPause()
	submit.MustClick()

	// Wait for result
	deadline := time.Now().Add(cfg.Timeout)

	for time.Now().Before(deadline) {
		// Check success
		if cfg.SuccessCheckJS != "" {
			ok := page.MustEval(cfg.SuccessCheckJS).Bool()
			if ok {
				return nil
			}
		}

		// Check failure
		if cfg.FailureCheckJS != "" {
			fail := page.MustEval(cfg.FailureCheckJS).Bool()
			if fail {
				return errors.New("login failed (credentials or checkpoint)")
			}
		}

		time.Sleep(500 * time.Millisecond)
	}

	return errors.New("login timeout")
}
