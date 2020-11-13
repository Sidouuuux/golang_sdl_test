package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	playerSpeed = 0.3
)

type player struct {
	tex  *sdl.Texture
	x, y float64
}

func newPlayer(renderer *sdl.Renderer, filename string, position [2]int) (p player, err error) {
	//chargement de l'image
	img, err := sdl.LoadBMP(filename)

	if err != nil {
		return player{}, fmt.Errorf("création ouverture fichier joueur : %v", err)
	}
	defer img.Free()

	//chargement d'une texture pour l'image
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("création texture joueur : %v", err)
	}
	p.x = float64(position[0])
	p.y = float64(position[1])
	return p, nil

}

func (p *player) DrawPlayer(renderer *sdl.Renderer) {
	//copie de la texture dans le renderer
	//			  //texture à copier
	renderer.Copy(p.tex,
		//src : partie de la texture à afficher
		&sdl.Rect{X: 0, Y: 0, W: 32, H: 32},
		//dst :  rectangle sur la fenetre où sera affiché la texture
		&sdl.Rect{X: int32(p.x), Y: int32(p.y), W: textureSize, H: textureSize})
}

func movePlayers(p1 *player, p2 *player) {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_RIGHT] == 1 && p1.x < screenWidth-textureSize {
		p1.x += playerSpeed
	}
	if keys[sdl.SCANCODE_LEFT] == 1 && p1.x != 0 && p1.x > 0 {
		p1.x -= playerSpeed
	}
	if keys[sdl.SCANCODE_UP] == 1 && p1.y != 0 && p1.y > 0 {
		p1.y -= playerSpeed
	}
	if keys[sdl.SCANCODE_DOWN] == 1 && p1.y < screenHeight-textureSize {
		p1.y += playerSpeed
	}

	if keys[sdl.SCANCODE_D] == 1 && p2.x < screenWidth-textureSize {
		p2.x += playerSpeed
	}
	if keys[sdl.SCANCODE_A] == 1 && p2.x != 0 && p2.x > 0 {
		p2.x -= playerSpeed
	}
	if keys[sdl.SCANCODE_W] == 1 && p2.y != 0 && p2.y > 0 {
		p2.y -= playerSpeed
	}
	if keys[sdl.SCANCODE_S] == 1 && p2.y < screenHeight-textureSize {
		p2.y += playerSpeed
	}
}
