package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	pc "github.com/mhbardsley/practice-choice"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// parse flags
	flags := pc.GetFlags()

	// load the input
	practiceParts := pc.LoadInput(flags.FilePtr)

	// prompt for values to add
	var name string
	var again string

	for strings.ToLower(again) != "n" {
		name = selected(practiceParts, scanner)
		delete(practiceParts, name)
		pc.WriteOutput(flags.FilePtr, practiceParts)
		fmt.Printf("Continue [Y/n]: ")
		scanner.Scan()
		again = scanner.Text()
	}
}

func selected(pp map[string]float64, scanner *bufio.Scanner) string {
	var choice string
	var choiceIndex int
	ppSlice := make([]string, len(pp))
	count := 0
	for key := range pp {
		ppSlice[count] = key
		count++
	}
LOOP:
	for i := range ppSlice {
		fmt.Printf("%d. %s", i, ppSlice[i])
		fmt.Println()
	}
	fmt.Printf("Please make a choice: ")
	scanner.Scan()
	choice = scanner.Text()
	choiceIndex, err := strconv.Atoi(choice)
	if err != nil || choiceIndex < 0 || choiceIndex >= len(ppSlice) {
		fmt.Println("Bad value entered. Please try again")
		goto LOOP
	}
	return ppSlice[choiceIndex]
}
