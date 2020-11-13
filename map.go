package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/veandco/go-sdl2/sdl"
)

type Level struct {
	Map [][]string
}

func LoadLevelFile(filename string) *Level {

	filebuffer, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("ouverture du fichier:", err)
		os.Exit(1)
	}

	inputdata := string(filebuffer)
	data := bufio.NewScanner(strings.NewReader(inputdata))

	level := &Level{}

	var text []string
	longestRow := 0
	row := 0
	for data.Scan() {
		text = append(text, data.Text())
		if len(text[row]) > longestRow {
			longestRow = len(text[row])
		}
		row++
	}

	level.Map = make([][]string, row)
	for i := range level.Map {
		level.Map[i] = make([]string, longestRow)
	}

	// lines := 0
	for i := range level.Map {
		for j := range level.Map[i] {
			level.Map[i][j] = string(text[i][j])
		}
	}

	return level
}

func printMap(renderer *sdl.Renderer, level *Level, filename string) {
	//chargement de l'image
	img, err := sdl.LoadBMP(filename)

	if err != nil {
		fmt.Errorf("création ouverture fichier joueur : %v", err)
		return
	}
	defer img.Free()
	var tileTex *sdl.Texture
	tileTex, err = renderer.CreateTextureFromSurface(img)
	if err != nil {
		fmt.Errorf("création texture joueur : %v", err)
		return
	}
	var tileRect sdl.Rect
	pleinFonce := sdl.Rect{384, 32, 32, 32}  //plein fonc�
	pleinClaire := sdl.Rect{32, 32, 32, 32}  //plein claire
	bordHaut := sdl.Rect{32, 0, 32, 32}      //bord haut
	bordBas := sdl.Rect{32, 64, 32, 32}      //bord bas
	bordGauche := sdl.Rect{0, 32, 32, 32}    //bord gauche
	bordDroit := sdl.Rect{64, 32, 32, 32}    //bord droit
	coinHautGauche := sdl.Rect{0, 0, 32, 32} //coin haut gauche
	coinHautDroit := sdl.Rect{64, 0, 32, 32} //coin haut droit
	coinBasGauche := sdl.Rect{0, 64, 32, 32} //coin haut gauche
	coinBasDroit := sdl.Rect{64, 64, 32, 32} //coin haut droit

	fmt.Println(level.Map)
	for i := range level.Map {
		fmt.Println("brrrrr")
		for j := range level.Map[i] {
			tileRect.X = int32(j * textureSize)
			tileRect.Y = int32(i * textureSize)

			fmt.Println(level.Map[i][j])
			switch level.Map[i][j] {
			case "A":
				renderer.Copy(tileTex, &pleinFonce, &tileRect)
				break

			case "B":
				renderer.Copy(tileTex, &pleinClaire, &tileRect)
				break

			case "C":
				renderer.Copy(tileTex, &bordHaut, &tileRect)
				break

			case "D":
				renderer.Copy(tileTex, &bordBas, &tileRect)
				break

			case "E":
				renderer.Copy(tileTex, &bordGauche, &tileRect)
				break

			case "F":
				renderer.Copy(tileTex, &bordDroit, &tileRect)
				break

			case "G":
				renderer.Copy(tileTex, &coinHautGauche, &tileRect)
				break

			case "H":
				renderer.Copy(tileTex, &coinHautDroit, &tileRect)
				break

			case "J":
				renderer.Copy(tileTex, &coinBasGauche, &tileRect)
				break

			case "K":
				renderer.Copy(tileTex, &coinBasDroit, &tileRect)
				break
			}
		}
	}
}
