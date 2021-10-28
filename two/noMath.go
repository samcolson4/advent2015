package noMath

import (
	"fmt"
	"strconv"
	"strings"
)

// func Calculate(s string) (int, error) {
// 	var output int

// 	return output, nil
// }

func StripText(s string) ([]int, error) {
	var measureInt []int
	measure := strings.Split(s, "x")

	for _, value := range measure {
		number, err := strconv.Atoi(value)

		if err != nil {
			fmt.Println(err)
		}

		measureInt = append(measureInt, number)
	}

	return measureInt, nil
}

func SmallestSide(input []int) (int, error) {

	length := input[0]
	width := input[1]
	height := input[2]

	a := length * height
	b := length * width
	c := height * width

	areas := []int{a, b, c}

	smallest := areas[0]
	for _, num := range areas[1:] {
		if num < smallest {
			smallest = num
		}
	}

	return smallest, nil
}
