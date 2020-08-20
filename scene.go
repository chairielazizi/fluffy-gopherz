package main

import (
	"context"
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"time"
)

type scene struct {
	time int

	bg *sdl.Texture
	birds []*sdl.Texture // slice of texture
}

func newScene(r *sdl.Renderer) (*scene,error) {
	bg, err := img.LoadTexture(r,"res/imgs/bg.jpg")
	if err != nil {
		fmt.Errorf("could not load background image: %v",err)
	}

	// animate the bird
	var birds []*sdl.Texture
	for i:=1; i<=4; i++ {
		path := fmt.Sprintf("res/imgs/frame-%d.png",i)
		bird, err := img.LoadTexture(r,path)
		if err != nil {
			fmt.Errorf("could not load background image: %v",err)
		}
		birds = append(birds,bird)
	}

	// return all where bg is texture, otherwise nil
	return &scene{bg:bg, birds: birds}, nil
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
	s.time++

	r.Clear()

	//background
	if err := r.Copy(s.bg,nil,nil); err != nil {
		return fmt.Errorf("could not copy background: %v",err)
	}

	// bird
	rect := &sdl.Rect{X: 10,Y: 300 - 43/2, W: 50, H: 43}
	i := s.time/10 % len(s.birds)
	if err := r.Copy(s.birds[i],nil,rect); err != nil {
		return fmt.Errorf("could not copy  the bird background: %v",err)
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
}
