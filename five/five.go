package five

import (
	"strings"
)

func CheckString(input string) (output bool) {
	if CheckVowels(input) && CheckDoubles(input) && CheckBanned(input) {
		output = true
	}

	return output
}

func CheckVowels(input string) (output bool) {
	var vowelCount int

	vowels := []string{"a", "e", "i", "o", "u"}

	splitString := strings.Split(input, "")

	for _, letter := range splitString {
		for _, vowel := range vowels {
			if letter == vowel {
				vowelCount++
			}
		}
	}

	if vowelCount >= 3 {
		output = true
	}

	return output

}

func CheckDoubles(input string) (output bool) {
	splitString := strings.Split(input, "")
	var doubleCount int

	for i, letter := range splitString {
		if i == len(splitString)-1 {
			continue
		} else if letter == splitString[i+1] {
			doubleCount++
		}
	}

	if doubleCount > 0 {
		output = true
	}

	return output
}

func CheckBanned(input string) (output bool) {
	var bannedCount int

	bannedStrings := []string{"ab", "cd", "pq", "xy"}
	splitString := strings.Split(input, "")

	for i, letter := range splitString {
		if i == len(splitString)-1 {
			continue
		}

		newString := letter + splitString[i+1]

		for _, banned := range bannedStrings {
			if newString == banned {
				bannedCount += 1
			} else {
				continue
			}
		}
	}

	if bannedCount > 0 {
		output = false
	} else {
		output = true
	}

	return output
}

func TwoCheckPairs(input string) (output bool) {
	var pairCount int
	splitString := strings.Split(input, "")

	for i, letter := range splitString {
		if i == len(splitString)-1 {
			continue
		}

		if letter == splitString[i+1] && i == 0 {
			pairCount++
		}

		if i != 0 && splitString[i+1] != splitString[i-1] && letter == splitString[i+1] {
			pairCount++
		}
	}

	if pairCount >= 2 {
		output = true
	}

	return output

}

func TwoCheckRepeats(input string) (output bool) {
	var doubleCount int
	splitString := strings.Split(input, "")

	for i, letter := range splitString {
		// checks if the index is the final one in the slice or the penultimate one.
		if i == len(splitString)-1 || i == len(splitString)-2 {
			continue
		}

		if letter == splitString[i+2] {
			doubleCount++
		}

	}

	if doubleCount > 0 {
		output = true
	}

	return output
}
