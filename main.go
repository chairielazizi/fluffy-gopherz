package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"runtime"
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

	//return drawTitle(r)
	if err := drawTitle(r,"Fluffy Gopherz"); err != nil {
		return fmt.Errorf("Could not draw title: %v",err)
	}

	//delay to destroy
	time.Sleep(1 *  time.Second)

	// draw background
	//img.Init()
	s, err := newScene(r)
	//if err := drawBackground(r); err != nil {
	//	fmt.Errorf("Could not draw background: %v",err)
	//}
	if err != nil {
		fmt.Errorf("Could not create scene: %v",err)
	}
	defer s.destroy()

	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()

	//go func() {
	//	time.Sleep(5 * time.Second)
	//	cancel()
	//}()
	// the same thing
	//time.AfterFunc(5 * time.Second, cancel)

	// goroutine continuously call the wait events, send it to the channel events, to the run func
	events := make(chan sdl.Event)
	errc := s.run(events,r)
	runtime.LockOSThread()
		for {
			select {
			case events <- sdl.WaitEvent():
			case err := <- errc:
				return err

			}
		}


	// wait until run() return
	//return <-

	//s.run(ctx,r)
	//if err := s.paint(r); err != nil {
	//	fmt.Errorf("could not paint the scene: %v",err)
	//}
	//
	//time.Sleep(5 * time.Second)
	//
	//return nil
}

func drawTitle(r *sdl.Renderer, text string) error {
	r.Clear() // to clear the buffer and present it
	// clear buffer,paint on it, and put it on the other side

	f, err := ttf.OpenFont("res/fonts/animal.ttf",170)
	if err != nil {
		return fmt.Errorf("Could not load font: %v",err)
	}
	// close the font used
	defer f.Close()

	c := sdl.Color{R: 255, G: 100, B: 0, A: 255}
	s, err := f.RenderUTF8Solid(text, c)
	if err != nil {
		return fmt.Errorf("Could not render title: %v",err)
	}
	// free the surface
	defer s.Free()

	// create texture from surface
	t, err := r.CreateTextureFromSurface(s)
	if err != nil {
		return fmt.Errorf("Could not create texture: %v",err)
	}
	defer t.Destroy()

	if err := r.Copy(t,nil,nil); err != nil {
		return fmt.Errorf("Could not copy texture: %v",err)
	}

	r.Present()

	return nil
}

//func drawBackground(r *sdl.Renderer) error {
//	r.Clear()
//
//	t, err := img.LoadTexture(r,"res/imgs/bg.jpg")
//	if err != nil {
//		fmt.Errorf("could not load background image: %v",err)
//	}
//	defer t.Destroy()
//
//	if err := r.Copy(t,nil,nil); err != nil {
//		return fmt.Errorf("could not copy background: %v",err)
//	}
//
//	r.Present()
//	return nil
//}
