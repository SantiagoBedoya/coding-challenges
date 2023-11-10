package counter

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func Count(countBytes, countLines, countWords, countCharacters bool, content string) string {
	values := make([]string, 0)
	if countBytes {
		nBytes := fmt.Sprint(countContentBytes(content))
		values = append(values, nBytes)
	}

	if countLines {
		lines := fmt.Sprint(countContentLines(content))
		values = append(values, lines)
	}

	if countWords {
		words := fmt.Sprint(countContentWords(content))
		values = append(values, words)
	}

	if countCharacters {
		characters := fmt.Sprint(countContentCharacters(content))
		values = append(values, characters)
	}

	if !countBytes && !countLines && !countWords && !countCharacters {
		nBytes := fmt.Sprint(countContentBytes(content))
		lines := fmt.Sprint(countContentLines(content))
		words := fmt.Sprint(countContentWords(content))
		characters := fmt.Sprint(countContentCharacters(content))
		values = append(values, nBytes, lines, words, characters)
	}

	return strings.Join(values, "\t")
}

func countContentBytes(content string) int {
	return len([]byte(content))
}

func countContentLines(content string) int {
	return len(strings.Split(content, "\r\n"))
}

func countContentWords(content string) int {
	return len(strings.Fields(content))
}

func countContentCharacters(content string) int {
	return utf8.RuneCountInString(content)
}
