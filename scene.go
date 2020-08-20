package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"time"
)

type scene struct {
	//time int // move to bird.go

	bg *sdl.Texture
	bird *bird
	//birds []*sdl.Texture // slice of texture // move to bird.go


	// add pipe
	pipe *pipe
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

	p, err := newPipe(r)
	if err != nil {
		return nil, err
	}

	// return all where bg is texture, otherwise nil
	return &scene{bg:bg, bird: b, pipe: p}, nil
}

// run in different goroutine
func (s *scene) run(events <-chan sdl.Event, r *sdl.Renderer) <-chan error {
	errc := make(chan error)
	go func() {
		defer close(errc)
		tick := time.Tick(10 * time.Millisecond)  // 100 frames per second

		for {
			select {
			case e := <- events:
				if done := s.handleEvent(e); done {
					return
				}
				log.Printf("event: %T", e)
				//return
			case <- tick:
				s.update()
				if s.bird.isDead() {
					drawTitle(r,"Game Over")
					time.Sleep(time.Second)
					s.restart()
				}
				if err := s.paint(r); err != nil {
					errc <- err // send error to the channel
				}
			}
		}
	}()

	return errc
}

func (s *scene) handleEvent(event sdl.Event) bool{
	switch event.(type) {
	case *sdl.QuitEvent:
		return true
	case *sdl.MouseButtonEvent:
		// click to jump
		s.bird.jump()
		return false
	case *sdl.MouseMotionEvent, *sdl.WindowEvent:
		// for not print on the log
	default:
		log.Printf("unknown event %T", event)
	}
		return false
}

func (s *scene) update() {
	s.bird.update()
	s.pipe.update()
}

func (s *scene) restart(){
	s.bird.restart()
	s.pipe.restart()
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

	// pipe
	if err := s.pipe.paint(r); err != nil {
		return err
	}

	r.Present()
	return nil
}

func (s *scene) destroy() {
	s.bg.Destroy()
	s.bird.destroy()
	s.pipe.destroy()
}
