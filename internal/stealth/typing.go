package stealth

import (
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

// TypeHuman types text char-by-char with human-like delays
func TypeHuman(el *rod.Element, text string) {
	for _, ch := range text {
		// Type one character
		el.Input(string(ch))

		// Random delay between keystrokes (human rhythm)
		delay := rand.Intn(120) + 60 // 60–180 ms
		time.Sleep(time.Duration(delay) * time.Millisecond)

		// Occasional micro pause (thinking / hesitation)
		if rand.Intn(12) == 0 {
			MicroPause()
		}
	}
}
