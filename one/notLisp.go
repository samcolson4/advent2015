package notLisp

func Calculate(s string) (int, error) {
	output := 0

	if s == "(" {
		output += 1
	} else {
		output -= 1
	}

	return output, nil
}
