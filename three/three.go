package three

func SetCoords(input string) (output [][]int, err error) {
	var newCoords []int

	startCoords := []int{0, 0}
	output = append(output, startCoords)

	if input == ">" {
		updateCoords := []int{startCoords[1] + 1, startCoords[0]}
		newCoords = updateCoords
	}

	output = append(output, newCoords)

	return output, nil
}
