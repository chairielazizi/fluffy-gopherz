package main

import (
	"context"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	//time int // move to bird.go

	bg *sdl.Texture
	bird *bird
	//birds []*sdl.Texture // slice of texture // move to bird.go
}

func newScene(r *sdl.Renderer) (*scene,error) {
	bg, err := img.LoadTexture(r,"res/imgs/bg2.jpg")
	if err != nil {
		fmt.Errorf("could not load background image: %v",err)
	}

	b, err := newBird(r)
	if err != nil {
		return nil, err
	}

	// return all where bg is texture, otherwise nil
	return &scene{bg:bg, bird: b}, nil
}

// run in different goroutine
func (s *scene) run(ctx context.Context, r *sdl.Renderer) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		for range time.Tick(10 * time.Millisecond) { // 100 frames per second
			select {
			case <-ctx.Done():
				return
			default:
				if err := s.paint(r); err != nil {
					errc <- err // send error to the channel
				}
			}
		}
	}()

	return errc
}

// method
func (s *scene) paint(r *sdl.Renderer) error {
	s.bird.time++

	r.Clear()

	//background
	if err := r.Copy(s.bg,nil,nil); err != nil {
		return fmt.Errorf("could not copy background: %v",err)
	}

	// bird
	if err := s.bird.paint(r); err != nil {
		return err
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.bird.destroy()
}
