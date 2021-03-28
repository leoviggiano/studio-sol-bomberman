package main

import "fmt"

func main() {
	inputExample := `6 7 3
	.......
	..XO...
	....O..
	.X.....
	OOX....
	OO.....`

	grid := NewGrid(inputExample)
	for _, val := range grid.Grid {
		fmt.Println(val)
	}
}
