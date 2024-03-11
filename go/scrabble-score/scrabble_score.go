package scrabble

import (
	"fmt"
	"strings"
)
 

func getLetterValue(letter rune) int {
	c := strings.ToUpper(fmt.Sprintf("%c", letter))

	switch c {
	case "A", "E", "I", "O", "U", "L", "N", "R", "S", "T":
		return 1
  case "D", "G":
		return 2
	case "B", "C", "M", "P":
		return 3
	case "F", "H", "V", "W", "Y":
		return 4
	case "K":
		return 5
	case "J", "X":
		return 8
	case "Q", "Z":
		return 10
	}

	return 0
}

func Score(word string) int {
	var sum int = 0;
	for _, v := range word {
		sum += getLetterValue(v)
	}

	return sum
}

