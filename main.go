package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 1600 //50 tiles - colonnes
	screenHeight = 960  //30 tiles - lignes
	textureSize  = 48
)

func main() {
	var files = [...]string{"game/images/xmas_player1.bmp", "game/images/xmas_player2.bmp", "game/images/map1.bmp"}
	var playersPosition = [2][2]int{
		{0, 0},
		{(screenWidth - textureSize), (screenHeight - textureSize)},
	}
	//initialisation de sdl
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initialisation de sdl :", err)
		return
	}

	// création d'une fenetre
	window, err := sdl.CreateWindow(
		//		  //position de la fenetre en x et y : mis en undefined						   //accélération avec OpenGL
		"Sidoux", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("création d'une fenetre:", err)
		return
	}
	defer window.Destroy()

	//création du renderer pour modifier le contenue de la fenêtre
	//									//ou mettre le renderer, la position, accélération graphique
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("création du renderer:", err)
		return
	}
	defer renderer.Destroy()

	player1, err := newPlayer(renderer, files[0], playersPosition[0])
	if err != nil {
		fmt.Println("création du joueur:", err)
		return
	}

	player2, err := newPlayer(renderer, files[1], playersPosition[1])
	if err != nil {
		fmt.Println("création du joueur:", err)
		return
	}

	for {
		//get event, et boucle sur la pile d'event
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			//les event sdl sont des pointeurs
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(255, 246, 255, 255)
		//remplie la fenetre avec la couleur
		renderer.Clear()

		level := LoadLevelFile("game/map/level.map")

		movePlayers(&player1, &player2)
		printMap(renderer, level, files[2])
		player1.DrawPlayer(renderer)
		player2.DrawPlayer(renderer)
		//affiche le renderer
		renderer.Present()
	}

}
