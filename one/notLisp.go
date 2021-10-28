package notLisp

import (
	"fmt"
	"strings"
)

func Calculate(s string) (int, error) {
	output := 0
	var outputSlice []int
	stringSlice := strings.Split(s, "")

	for _, letter := range stringSlice {
		if letter == "(" {
			output += 1
		} else if letter == ")" {
			output -= 1
		}

		outputSlice = append(outputSlice, output)
	}

	for i, v := range outputSlice {
		if v == -1 {
			fmt.Println(i)
		}
	}

	return output, nil
}
