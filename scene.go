package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type scene struct {
	bg *sdl.Texture
}

func newScene(r *sdl.Renderer) (*scene,error) {
	t, err := img.LoadTexture(r,"res/imgs/bg.jpg")
	if err != nil {
		fmt.Errorf("could not load background image: %v",err)
	}


	// return all where bg is texture, otherwise nil
	return &scene{bg:t}, nil
}

// method
func (s *scene) paint(r *sdl.Renderer) error {
	r.Clear()

	if err := r.Copy(s.bg,nil,nil); err != nil {
		return fmt.Errorf("could not copy background: %v",err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
