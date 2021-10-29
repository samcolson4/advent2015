package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	noMath "github.com/samcolson4/advent2015/two"
)

func main() {
	file, err := os.Open("./two/online_input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var data []string

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	output, err := noMath.SumPaper(data)

	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	fmt.Println(output)
}
