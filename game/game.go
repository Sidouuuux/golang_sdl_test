package game

import (
	"bufio"
	"fmt"
	"os"
)

type Tile rune

const (
	Wall  Tile = '#'
	Door  Tile = '|'
	Floor Tile = '.'
)

type Level struct {
	Map [][]Tile
}

func LoadLevelFile(filename string) *Level {
	//"map/level.map"
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	return nil
}
