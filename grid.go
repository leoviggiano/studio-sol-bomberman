package main

import (
	"errors"
	"log"
	"strings"

	c "github.com/theviggo/studio-sol-bomber-man/constants"
	utils "github.com/theviggo/studio-sol-bomber-man/utils"
)

// Grid -> Grid[Row][Column]
type Grid struct {
	Rows    int
	Columns int
	Seconds int
	Grid    [][]string
}

func NewGrid(input string) *Grid {
	splitted := strings.Split(input, "\n")
	config, err := utils.SliceAtoi(strings.Split(splitted[0], " "))

	if err != nil {
		log.Fatalf("\n[Grid NewGrid] Error on convert config string slice to an integer slice: %v", err)
		return nil
	}

	rows, columns, seconds := config[0], config[1], config[2]

	switch {
	case rows < 1:
		err = errors.New("Rows lesser than 1")
	case rows > 200:
		err = errors.New("Rows higher than 200")
	case columns < 1:
		err = errors.New("Columns lesser than 1")
	case columns > 200:
		err = errors.New("Columns higher than 200")
	case seconds < 1:
		err = errors.New("Seconds lesser than 1")
	case seconds > 1e9:
		err = errors.New("Seconds higher than 10^9")
	}

	if err != nil {
		log.Fatalf("\n[Grid NewGrid] Condition not filled: %v", err)
		return nil
	}

	grid, err := createGrid(splitted[1:], rows, columns)

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

func (g *Grid) Result() [][]string {
	return nil
}

func (g *Grid) Explode(row, column int) {
	g.Grid[row][column] = "Z"

	// Detonate right side
	for i := column + 1; i < g.Columns; i += 1 {
		if g.Grid[row][i] == c.OBSTACLE {
			break
		}

		g.Grid[row][i] = "Z"
	}

	// Detonate left side
	for i := column - 1; i >= 0; i -= 1 {
		if g.Grid[row][i] == c.OBSTACLE {
			break
		}

		g.Grid[row][i] = "Z"
	}

	// Detonate upper side
	for i := row - 1; i >= 0; i -= 1 {
		if g.Grid[i][column] == c.OBSTACLE {
			break
		}

		g.Grid[i][column] = "Z"
	}

	// Detonate lower side
	for i := row + 1; i < g.Rows; i += 1 {
		if g.Grid[i][column] == c.OBSTACLE {
			break
		}

		g.Grid[i][column] = "Z"
	}
}

func createGrid(gridInput []string, rows, columns int) ([][]string, error) {
	grid := make([][]string, 0)
	for _, row := range gridInput {
		rowSlice := make([]string, 0)

		for _, item := range filterItems(row) {
			rowSlice = append(rowSlice, item)
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
