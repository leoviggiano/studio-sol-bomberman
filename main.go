package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := make([]string, 0)
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Grid(Press enter with no letters in terminal once it's done):\n")
	for {
		// Scans a line from Stdin(Console)
		scanner.Scan()
		// Holds the string that scanned
		text := scanner.Text()
		if len(text) != 0 {
			input = append(input, text)
		} else {
			break
		}
	}

	grid := NewGrid(input)
	grid.Result(true)
}
