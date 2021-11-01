package three

import "strings"

func SetCoords(input string) (output [][]int, err error) {
	startCoords := []int{0, 0}
	output = append(output, startCoords)

	directions := strings.Split(input, "")

	for _, val := range directions {
		lastX := output[len(output)-1][0]
		lastY := output[len(output)-1][1]

		if val == ">" {
			updateCoords := []int{lastX + 1, lastY}
			output = append(output, updateCoords)
		} else if val == "^" {
			updateCoords := []int{lastX, lastY + 1}
			output = append(output, updateCoords)
		} else if val == "v" {
			updateCoords := []int{lastX, lastY - 1}
			output = append(output, updateCoords)
		} else if val == "<" {
			updateCoords := []int{lastX - 1, lastY}
			output = append(output, updateCoords)
		}
	}

	return output, nil
}
