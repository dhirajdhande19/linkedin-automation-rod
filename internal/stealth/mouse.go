package stealth

import (
	"math"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func MoveMouseHuman(page *rod.Page, toX, toY float64) {
	mouse := page.Mouse

	startX := rand.Float64()*200 + 100
	startY := rand.Float64()*200 + 100

	steps := rand.Intn(20) + 30 // 30–50 steps

	for i := 0; i <= steps; i++ {
		t := float64(i)/float64(steps)

		// Linear interpolation + tiny randomness
		x := lerp(startX, toX, t) + rand.Float64()*2
		y := lerp(startY, toY, t) + rand.Float64()*2

		mouse.MoveTo(proto.Point{X: x, Y: y})
		time.Sleep(time.Duration(rand.Intn(8)+5) * time.Millisecond)
	}
}


func lerp(a, b, t float64) float64 {
	return a + (b-a)*math.Min(math.Max(t, 0), 1)
}