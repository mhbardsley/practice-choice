package main

import (
	"fmt"

	pc "github.com/mhbardsley/practice-choice"
)

func main() {
	// parse flags
	flags := pc.GetFlags()

	// load the input
	practiceParts := pc.LoadInput(flags.FilePtr)

	fmt.Println(practiceParts)
}
