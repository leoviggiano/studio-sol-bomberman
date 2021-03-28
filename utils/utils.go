package utils

import (
	"strconv"
)

func SliceAtoi(sliceString []string) (sliceInt []int, err error) {
	for _, value := range sliceString {
		i, err := strconv.Atoi(value)
		if err != nil {
			return sliceInt, err
		}

		sliceInt = append(sliceInt, i)
	}
	return
}
