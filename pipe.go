package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

type pipes struct {
	mu sync.RWMutex

	texture *sdl.Texture
	speed int32

	pipes [] *pipe
}

func newPipes (r *sdl.Renderer) (*pipes, error) {
	texture, err := img.LoadTexture(r, "res/imgs/bottomPipe.png")
	if err != nil {
		fmt.Errorf("could not load pipe image: %v", err)
	}

	return &pipes{
		texture: texture,
		speed:   2,
	}, nil
}

func (ps *pipes) paint(r *sdl.Renderer) error {
	ps.mu.RLock()
	defer ps.mu.RUnlock()

	for _, p := range ps.pipes {
		if err := p.paint(r, ps.texture); err != nil {
			return err
		}
	}

	//rect := &sdl.Rect{X: p.x,Y: 600 - p.h, W: p.w, H: p.h}
	//flip := sdl.FLIP_NONE
	//// inverted pipe
	//if p.inverted {
	//	rect.Y = 0
	//	flip = sdl.FLIP_VERTICAL
	//}
	//if err := r.CopyEx(p.texture, nil, rect, 0, nil,flip); err != nil {
	//	return fmt.Errorf("could not copy background: %v",err)
	//}

	// bottom pipe
	//if err := r.Copy(p.texture,nil,rect); err != nil {
	//	return fmt.Errorf("could not copy background: %v",err)
	//}

	return nil
}

func (ps *pipes) touch(b *bird) {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, p := range ps.pipes {
		p.touch(b)
	}
}

func (ps *pipes) restart() {
	ps.mu.Lock()
	ps.mu.Unlock()

	ps.pipes = nil
}

func (ps *pipes) update() {
	ps.mu.RLock()
	defer ps.mu.RUnlock()
	for _, p := range ps.pipes{
		//p.update()
		p.mu.Lock()
		p.x -= ps.speed
		p.mu.Unlock()
	}
}

func (ps *pipes) destroy() {
	ps.mu.Lock()
	defer ps.mu.Unlock()

	ps.texture.Destroy()
}

type pipe struct {
	mu sync.RWMutex

	x int32
	h int32
	w int32
	inverted bool
}

// single pipe
func newPipe(r *sdl.Renderer) (*pipe, error) {
	//texture, err := img.LoadTexture(r,"res/imgs/bottomPipe.png")
	//if err != nil {
	//	fmt.Errorf("could not load pipe image: %v",err)
	//}

	return &pipe {
		//texture: texture,
		x: 400,
		h: 300,
		w: 50,
		//speed: 1,
		inverted: true,
	}, nil
}

func (p *pipe) touch(b *bird) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	b.touch(p)
}

func (p *pipe) paint(r *sdl.Renderer, texture *sdl.Texture) error {
	p.mu.RLock()
	defer p.mu.RUnlock()

	rect := &sdl.Rect{X: p.x,Y: 600 - p.h, W: p.w, H: p.h}
	flip := sdl.FLIP_NONE
	// inverted pipe
	if p.inverted {
		rect.Y = 0
		flip = sdl.FLIP_VERTICAL
	}
	if err := r.CopyEx(texture, nil, rect, 0, nil,flip); err != nil {
		return fmt.Errorf("could not copy background: %v",err)
	}

	// bottom pipe
	//if err := r.Copy(p.texture,nil,rect); err != nil {
	//	return fmt.Errorf("could not copy background: %v",err)
	//}

	return nil
}

func (p *pipe) restart() {
	p.mu.Lock()
	p.mu.Unlock()

	p.x = 400
}

//func (p *pipe) update() {
//	p.mu.Lock()
//	p.mu.Unlock()
//
//	p.x -= p.speed
//}
//
//func (p *pipe) destroy() {
//	p.mu.Lock()
//	p.mu.Unlock()
//
//	p.texture.Destroy()
//}