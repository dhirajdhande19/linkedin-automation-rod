package stealth

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// sleeps between min and max milis
func SleepRandom(minMs, maxMs int) {
	if minMs >= maxMs {
		time.Sleep(time.Duration(minMs) * time.Millisecond)
		return
	}

	delay := rand.Intn(maxMs-minMs) + minMs
	time.Sleep(time.Duration(delay) * time.Millisecond)
}

//  think like humans (short pause)
func Think() {
	SleepRandom(800, 1800)
}

// micro pause (tiny hesitation)
func MicroPause() {
	SleepRandom(120, 350)
}