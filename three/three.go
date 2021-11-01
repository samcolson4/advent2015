package three

func SetCoords(input string) (output [][]int, err error) {
	// var newCoords []int

	startCoords := []int{0, 0}
	output = append(output, startCoords)

	if input == ">" {
		updateCoords := []int{startCoords[1] + 1, startCoords[0]}
		output = append(output, updateCoords)
	} else if input == "^" {
		updateCoords := []int{startCoords[0], startCoords[1] + 1}
		output = append(output, updateCoords)
	}

	return output, nil
}
