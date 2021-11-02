package five

import "strings"

func CheckString(input string) (output bool, err error) {

	return output, nil
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
