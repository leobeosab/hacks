package main

/*
 * Aren't there easier ways or better tools to do this?
 * ...
 * ... let me do me
 */

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Generate number list incremented by value")
		fmt.Println("numgen start end increment (optional) precision")
		fmt.Println("Ex: numgen 0 10 .01")
		fmt.Println("Default precision is 0 ie no decimals")
		os.Exit(1)
	}

	var start float64
	var end float64
	var increment float64
	var err error

	start, err = strconv.ParseFloat(args[0], 64)
	end, err = strconv.ParseFloat(args[1], 64)
	increment, err = strconv.ParseFloat(args[2], 64)

	if err != nil {
		log.Fatal(err)
	}

	var o string
	if len(args) < 4 {
		o = "%.0f\n"
	} else {
		o = "%." + args[3] + "f\n"
	}

	for c := start; c < end; c += increment {
		fmt.Printf(o, c)
	}
}
