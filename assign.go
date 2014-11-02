package main

import (
	"flag"
	"fmt"
	"os"
)

func performSort(fileName string) {

}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: myprog [inputfile]\n")
	flag.PrintDefaults()
	os.Exit(2)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Input file is missing")
		os.Exit(1)
	}
	fmt.Printf("opening %s\n", args[0])

	performSort(args[0])
}
