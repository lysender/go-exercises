package main

import (
	"fmt"
	"unicode/utf8"
)

func ExampleRunes() {
	fmt.Println("================================")
	fmt.Println("        examples -> runes")
	fmt.Println("================================")
	runesExample1()
}

func runesExample1() {
	const s = "สวัสดี"

	fmt.Println("length of string thai word: hello")
	fmt.Println("len:", len(s))

	fmt.Println("indexing into a string to access raw bytes")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}
	fmt.Println()

	fmt.Println("Some utf-8 code points can span multiple bytes")
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	fmt.Println("using range to iterate runes instead of individual bytes")
	for idex, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idex)
	}

	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	if r == 't' {
		fmt.Println("found tee")
	} else if r == 'ส' {
		fmt.Println("found so sua")
	}
}
