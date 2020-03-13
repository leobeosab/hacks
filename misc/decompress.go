package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
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

	d = StringToByteArray("73 82 79 78 21 3 32 11 3 40 0 3 41 20 3 44 10 4 46 3 13 3 49 7 3 50 4 9 52 1 2 5 6 17 18 19 4 54 8 15 5 55 9 12 14 3 57 16")

	// First 4 are garbage and just say IRON
	d = d[4:]

	length := int(d[5])
	result := make(map[int]rune)

	fmt.Println(length)
	for i := 0; i < length; i++ {
		result[i] = rune(32)
	}

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

func StringToByteArray(s string) []byte {
	a := strings.Fields(s)
	fmt.Printf("%v", a)
	bytes := make([]byte, len(a))
	for i, b := range a {
		y, _ := strconv.Atoi(b)
		bytes[i] = byte(y)
	}
	fmt.Printf("%v", bytes)
	return bytes
}

func PrintMap(m *map[int]rune) {
	l := len(*m)
	fmt.Println(l)
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
