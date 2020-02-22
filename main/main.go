package main

import (
	"flag"
	"fmt"
	"log"
)

var (
	f = flag.Bool("f", false, "ignoring letters case")
	u = flag.Bool("u", false, "deleting same lines")
	r = flag.Bool("r", false, "Sorting descending")
	o = flag.String("o", "", "output to file")
	n = flag.Bool("n", false, "Sorting numbers")
	k = flag.Int("k", -1, "Sorting by column")
)

type Options struct {
	fFlag bool
	uFlag bool
	rFlag bool
	oFlag string
	nFlag bool
	kFlag int
}

func parseFlags() (*Options, error) {
	flag.Parse()

	if len(flag.Args()) < 1 {
		return nil, fmt.Errorf("not enough arguments")
	}

	return &Options{
		fFlag: *f,
		uFlag: *u,
		rFlag: *r,
		oFlag: *o,
		nFlag: *n,
		kFlag: *k,
	}, nil
}

func main() {
	opt, err := parseFlags()
	if err != nil {
		log.Fatal(err.Error())
	}

	lines, err := readFile(flag.Arg(0))
	if err != nil {
		log.Fatal(err.Error())
	}

	lines, err = Sorting(lines, opt)
	if err != nil {
		log.Fatal(err.Error())
	}
}
