package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	pc "github.com/mhbardsley/practice-choice"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	scanner := bufio.NewScanner(os.Stdin)
	// parse flags
	flags := pc.GetFlags()

	// load the input
	practiceParts := pc.LoadInput(flags.FilePtr)

	// prompt for values to add
	var name string
	var difficulty string
	var actualDifficulty float64
	var previousDifficulty float64
	var again string

	for strings.ToLower(again) != "n" {
		name = selected(practiceParts, scanner)
		fmt.Printf("New Difficulty: ")
		scanner.Scan()
		difficulty = scanner.Text()
		actualDifficulty, _ = strconv.ParseFloat(difficulty, 64)
		previousDifficulty = practiceParts[name]
		practiceParts[name] = new(previousDifficulty, actualDifficulty)
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

func new(prev float64, act float64) float64 {
	w := rand.Float64()
	return w*prev + (1-w)*act
}
