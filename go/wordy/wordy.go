package wordy

import (
	"strconv"
	"strings"
)

func Answer(question string) (int, bool) {
	words := strings.Fields(question)

	if strings.ToLower(words[0]) != "what" || strings.ToLower(words[1]) != "is" {
		return 0, false
	}

	words = words[2:]

	lastIndex := len(words) - 1
	lastWord := words[lastIndex]

	if lastWord[len(lastWord)-1] != '?' {
		return 0, false
	}

	words[lastIndex] = lastWord[:len(lastWord)-1]

	currentValue, err := strconv.Atoi(words[0])

	if err != nil {
		return 0, false
	}

	for i := 1; i < len(words); i++ {
		operator := words[i]		
		i++

		switch operator {
		case "multiplied", "divided":
			if i == len(words) || words[i] != "by" {
				return 0, false
			}
			i++
		}

		if i == len(words) {
			return 0, false
		}

		nextValue, err := strconv.Atoi(words[i])
		if err != nil {
			return 0, false
		}

		switch operator {
		case "plus":
			currentValue += nextValue
		case "minus": 
			currentValue -= nextValue
		case "multiplied":
			currentValue *= nextValue
		case "divided": 
			currentValue /= nextValue
		default:
			return 0, false
		}
	}

	return currentValue, true

}