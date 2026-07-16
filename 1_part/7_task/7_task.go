package main

import (
	"fmt"
	"unicode"
)

var vowelLetters = map[rune]struct{}{
	'а': {},
	'е': {},
	'ё': {},
	'и': {},
	'о': {},
	'у': {},
	'ы': {},
	'э': {},
	'ю': {},
	'я': {},
}

var consonantLetters = map[rune]struct{}{
	'б': {},
	'в': {},
	'г': {},
	'д': {},
	'ж': {},
	'з': {},
	'й': {},
	'к': {},
	'л': {},
	'м': {},
	'н': {},
	'п': {},
	'р': {},
	'с': {},
	'т': {},
	'ф': {},
	'х': {},
	'ц': {},
	'ч': {},
	'ш': {},
	'щ': {},
}

type StringStats struct {
	Vowel     int
	Consonant int
	Number    int
	Space     int
}

func main() {
	stats := countSymbols("Привет, мир! 123")

	fmt.Printf("%+v\n", stats)
}

func countSymbols(str string) StringStats {
	var stats StringStats

	for _, symbol := range str {
		symbol = unicode.ToLower(symbol)

		if symbol == ' ' {
			stats.Space++
			continue
		}

		if unicode.IsDigit(symbol) {
			stats.Number++
			continue
		}

		if _, ok := vowelLetters[symbol]; ok {
			stats.Vowel++
			continue
		}

		if _, ok := consonantLetters[symbol]; ok {
			stats.Consonant++
		}
	}

	return stats
}