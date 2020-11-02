package game

import (
	"bufio"
	"os"
)

// type Tile rune

// const (
// 	Wall  Tile = '#'
// 	Door  Tile = '|'
// 	Floor Tile = '.'
// )

type Level struct {
	Map [][]string
}

func LoadLevelFile(filename string) *Level {
	file, err := os.Open(filename)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	level := &Level{}

	var text []string
	longRow := 0
	col := 0

	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		text = append(text, scanner.Text())
		if len(text[col]) > longRow {
			longRow = len(text[col])
		}
		col++
	}

	for i := range level.Map {
		level.Map[i] = make([]string, longRow)

	}
	return level
}
