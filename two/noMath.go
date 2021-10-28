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
