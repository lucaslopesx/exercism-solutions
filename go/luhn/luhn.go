package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sum := 0
	isSecond := false

	for i := len(id) - 1; i >= 0; i--{
		x, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}

		if isSecond {
			x = x * 2
			if x > 9 {
				x -= 9
			}
		}
		
		sum += x
		isSecond = !isSecond
	}

	return sum % 10 == 0
}
