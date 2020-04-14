package main

import (
	"fmt"
	"github.com/Defernus/dwge"
	"math/rand"
)

var (
	menu_screen      *dwge.Image
	game_screen      *dwge.Image
	menu_scene_id    int
	game_scene_id    int
	current_scene_id int
	y, vy, timer     float64
	size             = 0.1
	x                = 0.2

	pipes            [4]float64
	high_score       = 0
	last_pipe, score int

	is_restart = false

	texture *dwge.Image
)

func onKeyPress(key string) {
	fmt.Printf("key [%v], scene %v\n", key, current_scene_id)
	if key == " " {
		if current_scene_id == game_scene_id {
			vy = .65
		} else {
			is_restart = true
		}
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

func create(win *dwge.MainWindow) error {
	w, h := win.GetSize()
	menu_scene_id = win.GetCurentSceneId()

	dwge.SetKeyDownEvent(onKeyPress)

	cw, ch := win.GetSize()

	menu_canvas, err := dwge.NewCanvas(0, 0, cw, ch, win)
	if err != nil {
		return err
	}

	if menu_screen, err = menu_canvas.GetImage(); err != nil {
		return err
	}

	game_scene_id = win.CreateNewScene()

	win.SetScene(game_scene_id)

	var game_canvas *dwge.Canvas
	game_canvas, err = dwge.NewCanvas(0, 0, cw, ch, win)
	if err != nil {
		return err
	}

	if game_screen, err = game_canvas.GetImage(); err != nil {
		return err
	}

	win.SetScene(menu_scene_id)

	button_img := dwge.NewImage(int(float64(w)*.1), int(float64(h)*.1))
	button_img.SetFillColor(0.5, 0.5, 0.5)
	button_img.Clear()
	button_img.SetFillColor(1, 1, 1)
	button_img.DrawRect(int(float64(w)*.025), int(float64(h)*.025), int(float64(w)*.05), int(float64(h)*.05))

	button_img_pressd := dwge.NewImage(int(float64(w)*.1), int(float64(h)*.1))
	button_img_pressd.SetFillColor(0, 0, 0)
	button_img_pressd.Clear()
	button_img_pressd.SetFillColor(.5, .5, .5)
	button_img_pressd.DrawRect(int(float64(w)*.025), int(float64(h)*.025), int(float64(w)*.05), int(float64(h)*.05))

	if _, err = dwge.NewButton(button_img, button_img_pressd, func() {
		win.SetScene(game_scene_id)
	}, int(float64(w)*.45), int(float64(h)*.5-float64(w)*.05), int(float64(w)*.1), int(float64(w)*.1), win); err != nil {
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

func menuLoop(win *dwge.MainWindow) error {
	w, h := func(a, b int) (float64, float64) { return float64(a), float64(b) }(win.GetSize())

	menu_screen.SetFillColor(.6, .9, 1)
	menu_screen.Clear()

	pixel_size := w * size
	menu_screen.DrawRotatedScaledImageAt(texture, int(w*x), int(y*h), int(pixel_size), int(pixel_size), .5, .5, vy)

	if is_restart {
		is_restart = false
		win.SetScene(game_scene_id)
		current_scene_id = game_scene_id
	}
	return nil
}

func gameLoop(win *dwge.MainWindow) error {
	w, h := func(a, b int) (float64, float64) { return float64(a), float64(b) }(win.GetSize())
	dt := dwge.GetDeltaTime()

	timer += dt

	vy -= dt
	y += vy * dt

	if (y < size/2) || (y > 1-size/2) {
		resetGame()
		win.SetScene(menu_scene_id)
		current_scene_id = menu_scene_id
	}

	lpx := 1 - timer/5 + float64(score)/2
	if x+.5*size > lpx && x-.5*size < lpx+.1 && (y-size*.5 < pipes[last_pipe]*.6 || y+size*.5 > pipes[last_pipe]*.6+.4) {
		resetGame()

		win.SetScene(menu_scene_id)
		current_scene_id = menu_scene_id
	}

	game_screen.SetFillColor(.6, .9, 1)
	game_screen.Clear()

	pixel_size := w * size
	game_screen.DrawRotatedScaledImageAt(texture, int(w*x), int(y*h), int(pixel_size), int(pixel_size), .5, .5, vy)

	game_screen.SetFillColor(0, .8, 0)
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
		game_screen.DrawRect(int(w*px), 0, int(w*.1), int(h*(pipes[i]*.6)))
		game_screen.DrawRect(int(w*px), int(h*(.4+pipes[i]*.6)), int(w*.1), int(h))
	}

	game_screen.SetFillColor(.1, .1, .1)
	game_screen.SetFontSize(20)

	game_screen.DrawText(fmt.Sprintf("High score: %v", high_score), 20, 60)
	game_screen.DrawText(fmt.Sprintf("Score: %v", score), 20, 40)
	game_screen.DrawText(fmt.Sprintf("FPS: %0.2f", 1./dt), 20, 20)

	return nil
}

func loop(win *dwge.MainWindow) error {
	switch scene_id := win.GetCurentSceneId(); scene_id {
	case menu_scene_id:
		return menuLoop(win)
	case game_scene_id:
		return gameLoop(win)
	default:
		return fmt.Errorf("wrong scene: %v", scene_id)
	}
}

func main() {
	resetGame()

	if err := dwge.Init(create, loop, 512, 512); err != nil {
		fmt.Printf("failed to init DWGE: %s\n", err)
	}
}
