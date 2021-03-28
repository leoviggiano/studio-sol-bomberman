package main

import (
	"errors"
	"log"
	"strings"

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
