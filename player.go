package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type player struct {
	tex *sdl.Texture
}

func newPlayer(renderer *sdl.Renderer) (p player, err error) {
	//chargement de l'image
	img, err := sdl.LoadBMP("game/images/tiles.bmp")

	if err != nil {
		return player{}, fmt.Errorf("création ouverture fichier joueur : %v", err)
	}
	defer img.Free()

	//chargement d'une texture pour l'image
	p.tex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		return player{}, fmt.Errorf("création texture joueur : %v", err)
	}
	return p, nil

}

func (p *player) DrawPlayer(renderer *sdl.Renderer) {
	//copie de la texture dans le renderer
	//			  //texture à copier
	renderer.Copy(p.tex,
		//src : partie de la texture à afficher
		&sdl.Rect{X: 32 * 8, Y: 0, W: 32, H: 32},
		//dst :  rectangle sur la fenetre où sera affiché la texture
		&sdl.Rect{X: 0, Y: 0, W: 100, H: 100})
}
