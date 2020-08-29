package main

import (
	"errors"
	"flag"
	"io/ioutil"
	"log"

	"github.com/p1ck4rd/famulus/pkg/cheatsheet"
	"github.com/p1ck4rd/famulus/pkg/parser"
)

var (
	input  = flag.String("i", "", "Input file")
	output = flag.String("o", "", "Output file")
)

func init() {
	flag.Parse()
}

func main() {
	if *input == "" {
		log.Fatal(errors.New("Input path must be specified"))
	}
	if *output == "" {
		log.Fatal(errors.New("Output path must be specified"))
	}

	content, err := ioutil.ReadFile(*input)
	if err != nil {
		log.Fatal(err)
	}

	res, err := parser.Unmarshal(content)
	if err != nil {
		log.Fatal(err)
	}

	err = cheatsheet.Generate(res, *output)
	if err != nil {
		log.Fatal(err)
	}
}
