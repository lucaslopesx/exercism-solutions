package parsinglogfiles

import (
	"fmt"
	"regexp"
)

func IsValidLine(text string) bool {
	re, err := regexp.Compile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)

	if err != nil {
		panic(err)
	}

	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re, err := regexp.Compile(`[<][*~=-]*[>]`)

	if err != nil {
		panic(err)
	}

	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re, err := regexp.Compile(`"[^"]*(?i)password[^"]*"`);

	if err != nil {
		panic(err)
	}

	count := 0
	for _, v := range lines {
		count += len(re.FindAllStringIndex(v, -1))
	}

	return count
}

func RemoveEndOfLineText(text string) string {
	re, err := regexp.Compile(`[0-9]*end-of-line[0-9]*`)

	if err != nil {
		panic(err)
	}

	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re, err := regexp.Compile(`User +((?i)[a-z0-9]+)`)

	if err != nil {
		panic(err)
	}

	var newLines []string

	for _, v := range lines {
		substring := re.FindStringSubmatch(v)

		if len(substring) == 0 {
			newLines = append(newLines, v)
			continue
		}
		
		if len(substring) > 0 {
			newLine := fmt.Sprintf("[USR] %s %s", substring[1], v)

			newLines = append(newLines, newLine)
		}
	}

	return newLines
}
