package main

import (
	"flag"
	"github.com/cabrel/stringsly"
	"log"
	"runtime"
)

var (
	fileToString     = flag.String("file", "", "File to generate strings for")
	parse            = flag.Bool("parse", false, "Parse the file directly instead of using system strings")
	includeLocations = flag.Bool("locations", false, "Include the location of the strings. Defaults to -td")
	currentOS        string
)

func main() {
	flag.Parse()

	if len(*fileToString) == 0 {
		log.Fatal("File must be provided")
	}

	if *parse {
		stringsly.ParseStrings(*fileToString)
	} else {
		if runtime.GOOS == "windows" {
			stringsly.RunStrings(*fileToString, *includeLocations, "")
		} else {
			stringsly.RunStrings(*fileToString, *includeLocations, "-eL")
			stringsly.RunStrings(*fileToString, *includeLocations, "-el")
			stringsly.RunStrings(*fileToString, *includeLocations, "-eB")
			stringsly.RunStrings(*fileToString, *includeLocations, "-eb")
			stringsly.RunStrings(*fileToString, *includeLocations, "-eS")
			stringsly.RunStrings(*fileToString, *includeLocations, "-es")
		}
	}
}
