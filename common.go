package pc

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path"
)

// Flags represents the range of flags that can be commonly passed to all commands
type Flags struct {
	FilePtr *string
}

// GetFlags gets the user-provided flags to the program
func GetFlags() Flags {
	var defaultFile string
	dirname, err := os.UserHomeDir()
	if err != nil {
		defaultFile = "input.json"
	} else {
		defaultFile = path.Join(dirname, ".practice-choice", "input.json")
	}
	filePtr := flag.String("f", defaultFile, "The input filename")

	flag.Parse()

	return Flags{filePtr}
}

// LoadInput loads a JSON file from the given file pointer
func LoadInput(filePtr *string) (data map[string]float64) {
	dataRaw, err := ioutil.ReadFile(*filePtr)
	if err != nil {
		log.Fatal("error opening file: ", err)
	}

	// unmarshall data into payload
	err = json.Unmarshal(dataRaw, &data)
	if err != nil {
		log.Fatal("error making sense of input file: ", err)
	}

	return data
}

// WriteOutput writes the JSON data to the JSON file
func WriteOutput(filePtr *string, data map[string]float64) {
	dataRaw, err := json.Marshal(data)
	if err != nil {
		log.Fatal("error making sense of data: ", err)
	}

	err = ioutil.WriteFile(*filePtr, dataRaw, 0644)
	if err != nil {
		log.Fatal("error writing to file: ", err)
	}
}
