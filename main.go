package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"time"
)

func main()  {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr,"%v",err)
		os.Exit(2)
	}
}

func run() error {
	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil{
		fmt.Errorf("Could not initialize SDL: %v",err)
		os.Exit(2)
	}
	defer sdl.Quit() //to quit the window

	if err := ttf.Init(); err != nil {
		return fmt.Errorf("Could not open window %v",err)
	}
	defer ttf.Quit()

	w,r,err := sdl.CreateWindowAndRenderer(800,600,sdl.WINDOW_SHOWN)
	if err != nil {
		return fmt.Errorf("Could not create window: %v",err)
	}
	defer w.Destroy()

	_ = r

	//delay to destroy
	time.Sleep(5 *  time.Second)

	return drawTitle(r)

	//return nil
}

func drawTitle(r *sdl.Renderer) error {
	f, err := ttf.OpenFont("res/fonts/patch.ttf",20)
	if err != nil {
		return fmt.Errorf("Could not load font: %v",err)
	}

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid("Fluffy Gopherz", c)
	if err != nil {
		return fmt.Errorf("Could not render title: %v",err)
	}

	// create texture from surface
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Could not create texture: %v",err)
	}

	if err := r.Copy(t,nil,nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v",err)
	}

	return nil
}
