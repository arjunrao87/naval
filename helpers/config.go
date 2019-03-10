package helpers

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// Read ...
func Read() []string {
	var quotes []string
	absPath, _ := filepath.Abs("../naval/resources/naval.json")
	jsonFile, err := os.Open(absPath)
	if err != nil {
		fmt.Println(err)
	}
	jsonParser := json.NewDecoder(jsonFile)
	jsonParser.Decode(&quotes)
	defer jsonFile.Close()
	return quotes
}
