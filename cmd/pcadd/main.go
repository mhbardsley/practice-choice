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
	var difficulty string
	var actualDifficulty float64
	var again string

	for strings.ToLower(again) != "n" {
		fmt.Printf("Name: ")
		scanner.Scan()
		name = scanner.Text()
		fmt.Printf("Difficulty: ")
		scanner.Scan()
		difficulty = scanner.Text()
		actualDifficulty, _ = strconv.ParseFloat(difficulty, 64)
		practiceParts[name] = actualDifficulty
		pc.WriteOutput(flags.FilePtr, practiceParts)
		fmt.Printf("Continue [Y/n]: ")
		scanner.Scan()
		again = scanner.Text()
	}
}
