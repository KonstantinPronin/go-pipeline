package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(flag.Args()) < 1 {
		log.Fatal("Not enough arguments")
	}

	expression := os.Args[1]
	fmt.Println(expression)
}
