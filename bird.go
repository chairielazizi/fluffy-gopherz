package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

const (
	gravity = 0.25
	jumpSpeed = -5
)

type bird struct {
	mu sync.RWMutex

	time int
	textures []*sdl.Texture

	// acceleration
	y, speed float64
	dead bool
}

func newBird(r *sdl.Renderer) (*bird, error) {
	// animate the bird
	var textures []*sdl.Texture
	for i:=1; i<=4; i++ {
		path := fmt.Sprintf("res/imgs/frame-%d.png",i)
		texture, err := img.LoadTexture(r,path)
		if err != nil {
			fmt.Errorf("could not load background image: %v",err)
		}
		textures = append(textures, texture)
	}
	return &bird{textures: textures, y: 300}, nil
}

func (b *bird) update() {
	b.mu.RLock()
	defer b.mu.RUnlock()

	// bird
	b.time++
	// add gravity
	b.y -= b.speed
	if b.y < 0 {
		// make it bounce when falling to the ground
		//b.speed = -b.speed
		//b.y = 0

		// the bird dead
		b.dead = true
	}
	b.speed += gravity
}

func (b *bird) paint(r *sdl.Renderer) error {

	rect := &sdl.Rect{X: 10,Y: (600 - int32(b.y)) - 43/2, W: 50, H: 43}

	i := b.time/10 % len(b.textures)
	if err := r.Copy(b.textures[i],nil,rect); err != nil {
		return fmt.Errorf("could not copy  the bird background: %v",err)
	}
	return nil
}

func (b *bird) destroy(){
	b.mu.Lock()
	defer b.mu.Unlock()

	for _, t := range b.textures {
		t.Destroy()
	}
}

func (b *bird) isDead() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.dead
}

func (b *bird) restart() {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.y = 300
	b.speed = 0
	b.dead = false
}

func (b *bird) jump() {
	b.speed = jumpSpeed
}