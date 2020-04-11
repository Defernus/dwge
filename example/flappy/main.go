package main

import (
	"fmt"
	"github.com/Defernus/dwge"
	"math/rand"
)

var (
	canvas       *dwge.Canvas
	screen       *dwge.Image
	y, vy, timer float64
	size         = 0.1
	x            = 0.2

	pipes            [4]float64
	high_score       = 0
	last_pipe, score int

	texture *dwge.Image
)

func onKeyPress(key string) {
	//fmt.Printf("keykode: %s\n", key)
	if key == " " {
		//fmt.Println("jump!")
		vy = .65
	}
}

func resetGame() {
	timer = 0
	y = .5
	vy = 0

	score = 0

	for i := range pipes {
		pipes[i] = rand.Float64()
	}

	last_pipe = 0
}

func create(win dwge.GuiElement) error {
	dwge.SetKeyDownEvent(onKeyPress)

	cwin := win.(dwge.GuiContainer)

	cw, ch := win.GetSize()
	var err error
	canvas, err = dwge.NewCanvas(0, 0, cw, ch, cwin)
	if err != nil {
		return err
	}
	screen, err = canvas.GetImage()
	if err != nil {
		return err
	}

	if texture, err = dwge.LoadImage("flappy.png"); err != nil {
		return err
	}

	return nil
}

func my_mod(a, b int) int {
	r := a % b
	if r < 0 {
		return b + a
	}
	return r
}

func loop(win dwge.GuiElement) error {
	w, h := func(a, b int) (float64, float64) { return float64(a), float64(b) }(canvas.GetSize())
	dt := dwge.GetDeltaTime()

	timer += dt

	vy -= dt
	y += vy * dt

	if (y < size/2) || (y > 1-size/2) {
		resetGame()
	}

	lpx := 1 - timer/5 + float64(score)/2
	if x+.5*size > lpx && x-.5*size < lpx+.1 && (y-size*.5 < pipes[last_pipe]*.6 || y+size*.5 > pipes[last_pipe]*.6+.4) {
		resetGame()
	}

	screen.SetFillColor(.6, .9, 1)
	screen.Clear()

	pixel_size := w * size
	screen.DrawRotatedScaledImageAt(texture, int(w*x), int(y*h), int(pixel_size), int(pixel_size), .5, .5, vy)

	screen.SetFillColor(0, .8, 0)
	for i := range pipes {
		px := 1 - timer/5 + float64(score+my_mod(i-last_pipe, 4))/2
		if px < -.2 {
			last_pipe = my_mod(last_pipe+1, 4)
			score++
			if score > high_score {
				high_score = score
			}
			pipes[i] = rand.Float64()
		}
		screen.DrawRect(int(w*px), 0, int(w*.1), int(h*(pipes[i]*.6)))
		screen.DrawRect(int(w*px), int(h*(.4+pipes[i]*.6)), int(w*.1), int(h))
	}

	screen.SetFillColor(.1, .1, .1)
	screen.SetFontSize(20)

	screen.DrawText(fmt.Sprintf("High score: %v", high_score), 20, 60)
	screen.DrawText(fmt.Sprintf("Score: %v", score), 20, 40)
	screen.DrawText(fmt.Sprintf("FPS: %0.2f", 1./dt), 20, 20)

	return nil
}

func main() {
	resetGame()

	if err := dwge.Init(create, loop, 512, 512); err != nil {
		fmt.Printf("failed to init DWGE: %s\n", err)
	}
}
