package main

import (
	"fmt"
	"math/rand"
	"time"

	pc "github.com/mhbardsley/practice-choice"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	// parse flags
	flags := pc.GetFlags()

	// load the input
	practiceParts := pc.LoadInput(flags.FilePtr)

	parts, weights := getWeights(practiceParts)

	var chosenIndex int
	r := rand.Float64() * weights[len(weights)-1]
	for i, weight := range weights {
		if r <= weight {
			chosenIndex = i
			break
		}
	}

	fmt.Printf("You should choose to practise %s", parts[chosenIndex])
	fmt.Println()
}

func getWeights(pp map[string]float64) (ppSlice []string, ppWeights []float64) {
	ppSlice = make([]string, len(pp))
	ppWeights = make([]float64, len(pp))
	count := 0
	for k, v := range pp {
		ppSlice[count] = k
		if count == 0 {
			ppWeights[count] = v
		} else {
			ppWeights[count] = v + ppWeights[count-1]
		}
		count++
	}
	return ppSlice, ppWeights
}
