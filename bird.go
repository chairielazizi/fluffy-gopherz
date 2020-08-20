package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type bird struct {
	time int
	textures []*sdl.Texture
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
	return &bird{textures: textures}, nil
}

func (b *bird) paint(r *sdl.Renderer) error {
	// bird
	b.time++

	rect := &sdl.Rect{X: 10,Y: 300 - 43/2, W: 50, H: 43}

	i := b.time/10 % len(b.textures)
	if err := r.Copy(b.textures[i],nil,rect); err != nil {
		return fmt.Errorf("could not copy  the bird background: %v",err)
	}
	return nil
}

func (b *bird) destroy(){
	for _, t := range b.textures {
		t.Destroy()
	}
}