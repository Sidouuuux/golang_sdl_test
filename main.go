package main

import "golang_sdl_test/game"

const (
	screenWidth  = 1600 //50 tiles - colonnes
	screenHeight = 960  //30 tiles - lignes
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

	// //chargement de l'image
	// img, err := sdl.LoadBMP("game/images/tiles.bmp")

	// if err != nil {
	// 	panic(err)
	// }
	// defer img.Free()

	// //chargement d'une texture pour l'image
	// mapTexture, err := renderer.CreateTextureFromSurface(img)
	// if err != nil {
	// 	panic(err)
	// }
	// defer mapTexture.Destroy()

	// for {
	// 	//get event, et boucle sur la pile d'event
	// 	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	// 		switch event.(type) {
	// 		//les event sdl sont des pointeurs
	// 		case *sdl.QuitEvent:
	// 			return
	// 		}
	// 	}
	// 	renderer.SetDrawColor(255, 246, 255, 255)
	// 	//remplie la fenetre avec la couleur
	// 	renderer.Clear()
	// 	//copie de la texture dans le renderer
	// 	//			  //texture à copier
	// 	renderer.Copy(mapTexture,
	// 		//src : partie de la texture à afficher
	// 		&sdl.Rect{X: 32 * 8, Y: 0, W: 32, H: 32},
	// 		//dst :  rectangle sur la fenetre où sera affiché la texture
	// 		&sdl.Rect{X: 0, Y: 0, W: 100, H: 100})
	// 	//affiche le renderer
	// 	renderer.Present()
	// }

}
