package main

import (
	"flag"
	"log"

	"famulus/gui"
	"famulus/internal/app/famulus"
)

var (
	input  = flag.String("i", "", "Input file")
	output = flag.String("o", "", "Output file")
)

func init() {
	flag.Parse()
}

func main() {
	if *input == "" && *output == "" {
		gui.Run()
	}

	if err := famulus.Generate(*input, *output); err != nil {
		log.Fatal(err)
	}
}
