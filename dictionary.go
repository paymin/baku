package baku

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

// the variable that holds all dictionary values
var notch map[byte]map[string]string

// defaultDirectory holds the dictionary folder with json alongside with this library as its default data source
const defaultDirectory = "/dictionary/*.json"

// InitializeDictionary or baku initiator starts with reading json files as its dictionary registered on memory runtime
// use this function inside your main package to populate dictionary to memory first before using the functions
func InitializeDictionary() error {
	paths, err := readAllFiles()
	if err != nil {
		return err
	}

	notch = make(map[byte]map[string]string)
	for _, path := range paths {
		err = populateDictionary(path)
		if err != nil {
			return err
		}
	}
	return nil
}

// readAllFiles returns all json files' path of the dictionary
func readAllFiles() ([]string, error) {
	_, filename, _, ok := runtime.Caller(1)
	if !ok {
		return nil, fmt.Errorf("Error#D11 - cannot get current dir")
	}

	// filename contains the path of this file as absolute path to go accessing the directory folder
	path := fmt.Sprint(filepath.Dir(filename), defaultDirectory)

	files, err := filepath.Glob(path)
	if err != nil {
		return nil, fmt.Errorf("Error#D12 - %s", err.Error())
	}
	return files, nil
}

// populateDictionary conducts the reading of dictionary json files and populate all values to notch variable
func populateDictionary(path string) error {
	dictionaryData, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("Error#D21 - %s", err.Error())
	}
	var dictionaryTemplate map[string][]string
	err = json.Unmarshal(dictionaryData, &dictionaryTemplate)
	if err != nil {
		return fmt.Errorf("Error#D22 - %s", err.Error())
	}

	for key, val := range dictionaryTemplate {
		for _, word := range val {
			if notch[[]byte(word[:1])[0]] == nil {
				notch[[]byte(word[:1])[0]] = make(map[string]string)
			}
			notch[[]byte(word[:1])[0]][word] = key
		}
	}
	return nil
}

// Formalize is the main process to search the inputed word and find the formal word as written on dictionary
func Formalize(word string) string {
	if notch == nil {
		return "dictionary has not been initialized"
	}
	bytedWord := []byte(strings.ToLower(word))
	if len(bytedWord) < 1 {
		return ""
	}

	notched := notch[bytedWord[0]]
	formalWord, ok := notched[word]
	if !ok {
		return fmt.Sprintf("kata %s belum tertulis di kamus baku", word)
	}
	return formalWord
}
