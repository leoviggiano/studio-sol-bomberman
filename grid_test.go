package main

import (
	"strings"
	"testing"
)

type IndexTest struct {
	input    string
	expected []string
}

func Equal(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

var oneSecondTests = []IndexTest{
	{
		input: `3 3 1
    .O.
    .X.
    ...`,
		expected: []string{
			".", "O", ".",
			".", "X", ".",
			".", ".", ".",
		},
	},
	{
		input: `5 5 1
    .O..X
    .X..O
    ....X
    XX.O.
    XOXO.`,
		expected: []string{
			".", "O", ".", ".", "X",
			".", "X", ".", ".", "O",
			".", ".", ".", ".", "X",
			"X", "X", ".", "O", ".",
			"X", "O", "X", "O", ".",
		},
	},
}

var twoSecondsTests = []IndexTest{
	{
		input: `3 3 2
    OOO
    OXO
    OOO`,
		expected: []string{
			"O", "O", "O",
			"O", "X", "O",
			"O", "O", "O",
		},
	},
	{
		input: `5 5 2
    OOOOX
    OXOOO
    OOOOX
    XXOOO
    XOXOO`,
		expected: []string{
			"O", "O", "O", "O", "X",
			"O", "X", "O", "O", "O",
			"O", "O", "O", "O", "X",
			"X", "X", "O", "O", "O",
			"X", "O", "X", "O", "O",
		},
	},
}

var threeSecondTests = []IndexTest{
	{
		input: `3 3 3
    .O.
    .X.
    ...`,
		expected: []string{
			".", ".", ".",
			"O", "X", "O",
			"O", "O", "O",
		},
	},
	{
		input: `5 5 3
    .O..X
    .X..O
    ....X
    XX.O.
    XOXO.`,
		expected: []string{
			".", ".", ".", ".", "X",
			"O", "X", ".", ".", ".",
			"O", "O", "O", ".", "X",
			"X", "X", ".", ".", ".",
			"X", ".", "X", ".", ".",
		},
	},
}

var nineHundredMillionsAndOneSecondTests = []IndexTest{
	{
		input: `3 3 900000001
    .O.
    .X.
    ...`,
		expected: []string{
			".", "O", ".",
			".", "X", ".",
			".", ".", ".",
		},
	},
	{
		input: `5 5 900000001
    .O..X
    .X..O
    ....X
    XX.O.
    XOXO.`,
		expected: []string{
			".", "O", ".", "O", "X",
			".", "X", ".", "O", "O",
			".", ".", ".", ".", "X",
			"X", "X", ".", "O", "O",
			"X", "O", "X", "O", "O",
		},
	},
}

var oneBillionSecondsTests = []IndexTest{
	{
		input: `3 3 1000000000
    .O.
    .X.
    ...`,
		expected: []string{
			".", ".", ".",
			"O", "X", "O",
			"O", "O", "O",
		},
	},
	{
		input: `5 5 1000000000
    .O..X
    .X..O
    ....X
    XX.O.
    XOXO.`,
		expected: []string{
			".", ".", ".", ".", "X",
			"O", "X", ".", ".", ".",
			"O", "O", "O", ".", "X",
			"X", "X", ".", ".", ".",
			"X", ".", "X", ".", ".",
		},
	},
}

func TestGridResult_OneSecond(t *testing.T) {
	for _, test := range oneSecondTests {
		output := NewGrid(strings.Split(test.input, "\n")).Result(false)
		expected := test.expected

		if !Equal(output, expected) {
			t.Errorf("Expected %v but got %v", expected, output)
		}
	}
}

func TestGridResult_TwoSeconds(t *testing.T) {
	for _, test := range twoSecondsTests {
		output := NewGrid(strings.Split(test.input, "\n")).Result(false)
		expected := test.expected

		if !Equal(output, expected) {
			t.Errorf("Expected %v but got %v", expected, output)
		}
	}
}

func TestGridResult_ThreeSeconds(t *testing.T) {
	for _, test := range threeSecondTests {
		output := NewGrid(strings.Split(test.input, "\n")).Result(false)
		expected := test.expected

		if !Equal(output, expected) {
			t.Errorf("Expected %v but got %v", expected, output)
		}
	}
}

func TestGridResult_NineHundredMillionsAndOneSeconds(t *testing.T) {
	for _, test := range nineHundredMillionsAndOneSecondTests {
		output := NewGrid(strings.Split(test.input, "\n")).Result(false)
		expected := test.expected

		if !Equal(output, expected) {
			t.Errorf("Expected %v but got %v", expected, output)
		}
	}
}

func TestGridResult_OneBillionSeconds(t *testing.T) {
	for _, test := range oneBillionSecondsTests {
		output := NewGrid(strings.Split(test.input, "\n")).Result(false)
		expected := test.expected

		if !Equal(output, expected) {
			t.Errorf("Expected %v but got %v", expected, output)
		}
	}
}
