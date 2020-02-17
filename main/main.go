package main

import (
	"flag"
	"log"
)

var f = flag.Bool("f", false, "ignoring letters case")
var u = flag.Bool("u", false, "deleting same lines")
var r = flag.Bool("r", false, "Sorting descending")
var o = flag.String("o", "", "output to file")
var n = flag.Bool("n", false, "Sorting numbers")
var k = flag.Int("k", -1, "Sorting by column")

func main() {
	flag.Parse()

	if len(flag.Args()) < 1 {
		log.Fatal("Not enough arguments")
	}

	lines, err := readFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err.Error())
	}

	lines, err = Sorting(lines)
	if err != nil {
		log.Fatal(err.Error())
	}
}
