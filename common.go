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

// PracticePart represents a single part that may be selected for practice
type PracticePart struct {
	Name       string  `json:"name"`
	Difficulty float64 `json:"difficulty"`
}

// GetFlags gets the user-provided flags to the program
func GetFlags() Flags {
	var defaultFile string
	dirname, err := os.UserHomeDir()
	if err != nil {
		defaultFile = "input.json"
	} else {
		defaultFile = path.Join(dirname, "input.json")
	}
	filePtr := flag.String("f", defaultFile, "The input filename")

	flag.Parse()

	return Flags{filePtr}
}

// LoadInput loads a JSON file from the given file pointer
func LoadInput(filePtr *string) (data []PracticePart) {
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
