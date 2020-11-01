package main

import (
	"golang_sdl_test/game"
)

const (
	screenWidth  = 800
	screenHeight = 800
)

func main() {
	game.LoadLevelFile("game/map/level.map")
	// //initialisation de sdl
	// if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
	// 	fmt.Println("initialisation de sdl :", err)
	// 	return
	// }

	// // création d'une fenetre
	// window, err := sdl.CreateWindow(
	// 	//		  //position de la fenetre en x et y : mis en undefined						   //accélération avec OpenGL
	// 	"Sidoux", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_OPENGL)
	// if err != nil {
	// 	fmt.Println("création d'une fenetre:", err)
	// 	return
	// }
	// defer window.Destroy()

	// //création du renderer pour modifier le contenue de la fenêtre
	// //									//ou mettre le renderer, la position, accélération graphique
	// renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	// if err != nil {
	// 	fmt.Println("création du renderer:", err)
	// 	return
	// }
	// defer renderer.Destroy()

	// for {
	// 	//get event, et boucle sur la pile d'event
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		case *sdl.QuitEvent:
	// 			return
	// 		}
	// 	}
	// 	renderer.SetDrawColor(255, 133, 255, 255)
	// 	//remplie la fenetre avec la couleur
	// 	renderer.Clear()
	// 	//affiche le renderer
	// 	renderer.Present()
	// }

}
