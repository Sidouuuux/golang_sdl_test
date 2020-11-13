package game

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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
