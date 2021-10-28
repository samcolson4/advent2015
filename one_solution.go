package main

import (
	"fmt"
	"io/ioutil"

	notLisp "github.com/samcolson4/advent2015/one"
)

func main() {
	data, err := ioutil.ReadFile("./one/online_input.txt")

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	output, err := notLisp.Calculate(string(data))

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println(output)
}
