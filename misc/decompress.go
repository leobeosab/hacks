package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/*
Description of compression
IRON-lenghtofstring-charactermap[offset-character-charpositions]
*/

func main() {
	var file string
	if len(os.Args) > 1 {
		file = os.Args[1]
	} else {
		file = "intercept3.iron"
	}
	d, e := ioutil.ReadFile(file)
	if e != nil {
		panic("fucked up on the file")
	}

	// First 4 are garbage and just say IRON
	d = d[4:]

	result := make(map[int]rune)

	for bytePos := 1; bytePos < len(d); {
		offset := int(d[bytePos])
		char := rune(d[bytePos+1])
		positions := d[bytePos+2 : bytePos+offset]

		AddCharToMap(&result, char, positions)

		fmt.Printf("Character: %s Positions -> ", string(char))
		PrintHex(positions)

		bytePos += offset
	}

	PrintMap(&result)
}

func PrintMap(m *map[int]rune) {
	l := len(*m)
	for i := 0; i < l; i++ {
		fmt.Printf("%c", (*m)[i])
	}
}

func AddCharToMap(m *map[int]rune, c rune, bs []byte) {
	for _, b := range bs {
		(*m)[int(b)] = c
	}
}

// Much more pretty than Go's default
func PrintHex(bs []byte) {
	for _, b := range bs {
		fmt.Printf("%.2X ", b)
	}
	fmt.Printf("\n")
}
