package notLisp

import (
	"strings"
)

func Calculate(s string) (int, error) {
	output := 0
	stringSlice := strings.Split(s, "")

	for _, letter := range stringSlice {
		if letter == "(" {
			output += 1
		} else if letter == ")" {
			output -= 1
		}
	}

	return output, nil
}
