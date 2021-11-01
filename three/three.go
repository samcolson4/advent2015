package three

import (
	"strings"
)

type coordinate struct {
	xCoord int
	yCoord int
}

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

func CalculateHouses(coords []coordinate) (output int, err error) {
	var beenThere []coordinate

uniqueLoop:
	for _, v := range coords {
		for i, u := range beenThere {
			if v.xCoord == u.xCoord && v.yCoord == u.yCoord {
				beenThere[i] = v
				continue uniqueLoop
			}
		}
		beenThere = append(beenThere, v)
	}

	output = len(beenThere)

	return output, nil
}

func MakeCoords(coords [][]int) (output []coordinate) {

	for _, coord := range coords {
		coStruct := coordinate{
			xCoord: coord[0],
			yCoord: coord[1],
		}
		output = append(output, coStruct)
	}

	return output
}

func SplitInput(input string) (SantaCoords string, RoboCoords string, err error) {
	var sCoord []string
	var rCoord []string

	directions := strings.Split(input, "")

	for index, element := range directions {
		if index%2 == 0 {
			sCoord = append(sCoord, element)
		} else {
			rCoord = append(rCoord, element)
		}
	}

	SantaCoords = strings.Join(sCoord[:], ",")
	RoboCoords = strings.Join(rCoord[:], ",")

	return SantaCoords, RoboCoords, nil
}
