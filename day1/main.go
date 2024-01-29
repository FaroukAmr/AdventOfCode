package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += getValues(scanner.Text())
	}

	fmt.Println(total)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func getValues(s string) int {
	firstNumber, secondNumber := 0, 0

	for _, v := range s {
		if isDigit(v) {
			firstNumber = int(v - '0')
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if isDigit(rune(s[i])) {
			secondNumber = int(rune(s[i]) - '0')
			break
		}
	}

	return firstNumber*10 + secondNumber
}

func isDigit(c rune) bool {
	return c >= '0' && c <= '9'
}
