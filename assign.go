package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"math/rand"
	"os"
)

func readCSVFile(fileName string) [][]string {
	csvfile, err := os.Open(fileName)

	if err != nil {
		fmt.Println(err)
		return nil
	}

	defer csvfile.Close()
	reader := csv.NewReader(csvfile)
	reader.FieldsPerRecord = -1

	rawCSVdata, err := reader.ReadAll()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return rawCSVdata
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func performSort(filename string) {
	rawCSVdata := readCSVFile(filename)
	for i := range rawCSVdata {
		j := rand.Intn(i + 1)
		rawCSVdata[i], rawCSVdata[j] = rawCSVdata[j], rawCSVdata[i]
	}

	selections := map[string]bool{}

	for _, row := range rawCSVdata {
		for j, val := range row {
			if j > 0 && (!selections[val]) {
				selections[val] = true
				fmt.Printf("%s,%s\n", row[0], val)
				break
			}
		}
	}
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
