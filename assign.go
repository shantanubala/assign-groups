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
	csvOut, err := os.Create("output.csv")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer csvOut.Close()
	writer := csv.NewWriter(csvOut)

	rawCSVdata := readCSVFile(filename)
	k := 0
	for i := range rawCSVdata {
		k = i
		if len(rawCSVdata[i]) == 0 || len(rawCSVdata[i][0]) == 0 {
			break
		}
	}

	allTeams := rawCSVdata[:k]
	allProjects := rawCSVdata[k+1:]

	fmt.Printf("Total Teams: %d\n", len(allTeams))
	fmt.Printf("Total Projects: %d\n", len(allProjects))
	writer.Write([]string{"Total Projects", fmt.Sprintf("%d", len(allTeams))})
	writer.Write([]string{"Total Teams", fmt.Sprintf("%d", len(allProjects))})

	for i := range allTeams {
		j := rand.Intn(i + 1)
		allTeams[i], allTeams[j] = allTeams[j], allTeams[i]
	}

	selections := map[string]bool{}
	teams := map[string]string{}
	pickPositions := map[int]int{}
	randomPicks := 0
	noPicks := 0

	totalPicks := (len(allTeams[0]) - 1) / 2

	for _, row := range allTeams {
		for j, val := range row {
			if j > 0 && j <= totalPicks && (!selections[val]) {
				pickPositions[j] = pickPositions[j] + 1
				selections[val] = true
				teams[row[0]] = val
				fmt.Printf("%s,%s,Pick #%d\n", row[0], val, j)
				writer.Write([]string{row[0], val, fmt.Sprintf("Pick #%d", j)})
				break
			}
		}
	}

	fmt.Print("\n")

	for _, row := range allTeams {
		_, exists := teams[row[0]]

		if !exists {
			for _, projects := range allProjects {
				project := projects[0]
				_, taken := selections[project]
				if taken {
					continue
				}

				selectProj := true

				for _, sel := range row {
					if project == sel {
						selectProj = false
						break
					}
				}

				if !selectProj {
					continue
				}

				fmt.Printf("%s,%s,Random Pick\n", row[0], project)
				writer.Write([]string{row[0], project, "Random Pick"})
				randomPicks += 1
				selections[project] = true
				teams[row[0]] = project
				break
			}
		}
	}

	fmt.Print("\n")

	for _, row := range allTeams {
		_, exists := teams[row[0]]

		if !exists {
			fmt.Printf("%s,None\n", row[0])
			writer.Write([]string{row[0], "None", "None"})
			noPicks += 1
		}
	}

	fmt.Print("\n")

	for k, v := range pickPositions {
		fmt.Printf("Pick #%d: %d\n", k, v)
		writer.Write([]string{fmt.Sprintf("Pick #%d", k), fmt.Sprintf("%d", v)})
	}
	fmt.Printf("Random Picks: %d\n", randomPicks)
	writer.Write([]string{"Random Picks", fmt.Sprintf("%d", randomPicks)})
	fmt.Printf("Unable to Pick: %d\n", noPicks)
	writer.Write([]string{"Unable to Pick", fmt.Sprintf("%d", noPicks)})

	writer.Flush()
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
	fmt.Printf("opening %s\n\n", args[0])

	performSort(args[0])
}
