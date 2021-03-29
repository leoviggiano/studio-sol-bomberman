package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	utils "github.com/theviggo/studio-sol-bomber-man/utils"
)

const (
	EMPTY    = "."
	BOMB     = "O"
	OBSTACLE = "X"
)

// Grid -> Grid[Row][Column]
type Grid struct {
	Rows           int
	Columns        int
	Seconds        int
	Grid           [][]Tile
	FirstExplosion [][]string
	LastExplosion  [][]string
}

type Tile struct {
	Item    string
	Seconds int
	Row     int
	Column  int
}

func NewGrid(input []string) *Grid {
	config, err := utils.SliceAtoi(strings.Split(input[0], " "))

	if err != nil {
		log.Fatalf("\n[Grid NewGrid] Error on convert config string slice to an integer slice: %v", err)
		return nil
	}

	rows, columns, seconds := config[0], config[1], config[2]

	switch {
	case rows < 1:
		err = errors.New("Rows less than 1")
	case rows > 200:
		err = errors.New("Rows greater than 200")
	case columns < 1:
		err = errors.New("Columns less than 1")
	case columns > 200:
		err = errors.New("Columns greater than 200")
	case seconds < 1:
		err = errors.New("Seconds less than 1")
	case seconds > 1e9:
		err = errors.New("Seconds greater than 10^9")
	}

	if err != nil {
		log.Fatalf("\n[Grid NewGrid] Condition not filled: %v", err)
		return nil
	}

	grid, err := createGrid(input[1:], rows, columns)

	if err != nil {
		log.Fatalf("\n[Grid CreateGrid] Error on create grid: %v", err)
		return nil
	}

	return &Grid{
		Rows:    rows,
		Columns: columns,
		Seconds: seconds,
		Grid:    grid,
	}
}

func (g *Grid) Result(printResult bool) []string {
	// After 5 iterations the result will be the same, depending if seconds is odd or even
	for i := 1; i <= g.Seconds && i <= 5; i += 1 {
		g.addSecond()
		switch i {
		case 2:
			g.fillGrid()

		case 4:
			g.fillGrid()
		}
	}

	output := make([]string, 0)
	result := ""

	switch {
	case g.Seconds > 4:
		grid := make([][]string, 0)
		// Verify if it's the first pattern of explosion, or if it's the second one
		isFirstExplosion := g.Seconds%2 == 0

		if isFirstExplosion {
			grid = g.FirstExplosion
		} else {
			grid = g.LastExplosion
		}

		for _, row := range grid {
			for _, item := range row {
				result += fmt.Sprintf(" %v ", item)
				output = append(output, item)
			}

			result += "\n"
		}

	default:
		for _, row := range g.Grid {
			for _, item := range row {
				result += fmt.Sprintf(" %v ", item.Item)
				output = append(output, item.Item)
			}

			result += "\n"
		}
	}

	if printResult {
		fmt.Println(result)
	}

	return output
}

func (g *Grid) Explode(row, column int) {
	g.Grid[row][column] = Tile{
		Item:    EMPTY,
		Seconds: 0,
		Row:     row,
		Column:  column,
	}

	// Detonate right side
	for i := column + 1; i < g.Columns; i += 1 {
		if g.Grid[row][i].Item == OBSTACLE {
			break
		}

		g.Grid[row][i] = Tile{
			Item:    EMPTY,
			Seconds: 0,
			Row:     row,
			Column:  i,
		}
	}

	// Detonate left side
	for i := column - 1; i >= 0; i -= 1 {
		if g.Grid[row][i].Item == OBSTACLE {
			break
		}

		g.Grid[row][i] = Tile{
			Item:    EMPTY,
			Seconds: 0,
			Row:     row,
			Column:  i,
		}
	}

	// Detonate upper side
	for i := row - 1; i >= 0; i -= 1 {
		if g.Grid[i][column].Item == OBSTACLE {
			break
		}

		g.Grid[i][column] = Tile{
			Item:    EMPTY,
			Seconds: 0,
			Row:     i,
			Column:  column,
		}
	}

	// Detonate lower side
	for i := row + 1; i < g.Rows; i += 1 {
		if g.Grid[i][column].Item == OBSTACLE {
			break
		}

		g.Grid[i][column] = Tile{
			Item:    EMPTY,
			Seconds: 0,
			Row:     i,
			Column:  column,
		}
	}
}

func createGrid(gridInput []string, rows, columns int) ([][]Tile, error) {
	grid := make([][]Tile, 0)
	for rowIndex, row := range gridInput {
		rowSlice := make([]Tile, 0)

		for columnIndex, item := range filterItems(row) {
			rowSlice = append(rowSlice, Tile{
				Item:    item,
				Seconds: 0,
				Row:     rowIndex,
				Column:  columnIndex,
			})
		}

		if len(rowSlice) != columns {
			return nil, errors.New("Length of columns specified on first line does not match with input")
		}

		grid = append(grid, rowSlice)
	}

	if len(grid) != rows {
		return nil, errors.New("Length of rows specified on first line does not match with input")
	}

	return grid, nil
}

// fillGrid fills all the empty spaces on grid with bombs with 0 seconds
func (g *Grid) fillGrid() {
	for rowIndex, row := range g.Grid {
		for columnIndex, tile := range row {
			if tile.Item == EMPTY {
				g.Grid[rowIndex][columnIndex] = Tile{
					Item:    BOMB,
					Seconds: 0,
					Row:     rowIndex,
					Column:  columnIndex,
				}
			}
		}
	}
}

// addSeconds adds 1 second to all tiles on grid and save the current state if had an explosion
func (g *Grid) addSecond() {
	hadExplosion := false
	toExplode := make([]Tile, 0)

	for rowIndex, row := range g.Grid {
		for columnIndex, tile := range row {
			if tile.Item == BOMB && tile.Seconds == 2 {
				toExplode = append(toExplode, tile)
				hadExplosion = true
			} else {
				g.Grid[rowIndex][columnIndex].Seconds += 1
			}
		}
	}

	if hadExplosion {
		for _, tile := range toExplode {
			g.Explode(tile.Row, tile.Column)
		}

		if g.FirstExplosion == nil {
			g.FirstExplosion = g.copyArray()
		} else if g.LastExplosion == nil {
			g.LastExplosion = g.copyArray()
		}
	}
}

// copyArray copy the current state of grid to an array of strings
func (g *Grid) copyArray() [][]string {
	newArray := make([][]string, 0)

	for _, row := range g.Grid {
		newRow := make([]string, 0)
		for _, value := range row {
			newRow = append(newRow, value.Item)
		}
		newArray = append(newArray, newRow)
	}

	return newArray
}

// filterItems returns only allowed values to create a grid
func filterItems(row string) (newInput []string) {
	for _, item := range row {
		if isAllowed(item) {
			newInput = append(newInput, string(item))
		}
	}

	return
}

// isAllowed verify if the coming value is allowed or not
func isAllowed(item rune) bool {
	allowedValues := [3]string{".", "X", "O"}
	for _, value := range allowedValues {
		if string(item) == value {
			return true
		}
	}

	return false
}
