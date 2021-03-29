package main

func main() {
	inputExample := `6 7 7
	...O..O
	..X..O.
	...O...
	.X...O.
	..X....
	.O.....`

	grid := NewGrid(inputExample)
	grid.Result()
	// 	printAll := true

	// 	if printAll {
	// 		for _, arr := range grid.Grid {
	// 			fmt.Printf("[")
	// 			for _, val := range arr {
	// 				fmt.Printf(" %v ", val.Item)
	// 			}
	// 			fmt.Printf("]\n")
	// 		}
	// 	} else {
	// 		for _, arr := range grid.Grid {
	// 			fmt.Printf("[")
	// 			for _, val := range arr {
	// 				fmt.Printf(" %v(%d) ", val.Item, val.Seconds)
	// 			}
	// 			fmt.Printf("]\n")
	// 		}
	// 	}
}
