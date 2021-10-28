package noMath

import (
	"fmt"
	"strconv"
	"strings"
)

func Calculate(input []int) (int, error) {
	var output int

	length := input[0]
	width := input[1]
	height := input[2]

	box := (2 * length * width) + (2 * width * height) + (2 * height * length)

	smallestSide, err := SmallestSide(input)

	if err != nil {
		fmt.Println(err)
	}

	output = box + smallestSide

	return output, nil
}

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

func SumPaper(input []string) (int, error) {
	var total int

	for _, present := range input {
		strippedValues, err := StripText(present)
		if err != nil {
			fmt.Println(err)
		}

		wrappingNeeded, err := Calculate(strippedValues)
		if err != nil {
			fmt.Println(err)
		}

		total += wrappingNeeded
	}

	return total, nil
}
