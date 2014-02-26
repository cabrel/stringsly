package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"
)

var (
	fileToString     = flag.String("file", "", "File to generate strings for")
	includeLocations = flag.Bool("locations", false, "Include the location of the strings. Defaults to -td")
	currentOS        string
)

func main() {
	flag.Parse()

	if len(*fileToString) == 0 {
		log.Fatal("File must be provided")
	}

	if runtime.GOOS == "windows" {
		runStrings(*fileToString, *includeLocations, "")
	} else {
		runStrings(*fileToString, *includeLocations, "-eL")
		runStrings(*fileToString, *includeLocations, "-el")
		runStrings(*fileToString, *includeLocations, "-eB")
		runStrings(*fileToString, *includeLocations, "-eb")
		runStrings(*fileToString, *includeLocations, "-eS")
		runStrings(*fileToString, *includeLocations, "-es")
	}
}

func combine(b1 []byte, b2 []byte) []byte {

	newslice := make([]byte, len(b1)+len(b2))
	copy(newslice, b1)
	copy(newslice[len(b1):], b2)
	return newslice
}

func runStrings(filePath string, incLoc bool, opts string) {
	var res []byte
	var argv []string
	var ext string

	if runtime.GOOS == "windows" {
		if incLoc {
			argv = append(argv, "-o")
		}
	} else {
		argv = append(argv, "-a")

		if incLoc {
			argv = append(argv, "-td")
		}

		argv = append(argv, opts)
	}

	argv = append(argv, filePath)

	cmd := exec.Command("strings", argv...)
	lres, _ := cmd.CombinedOutput()

	if opts == "-el" {
		ext = "16bitle"
	} else if opts == "-eL" {
		ext = "32bitle"
	} else if opts == "-eB" {
		ext = "32bitbe"
	} else if opts == "-eb" {
		ext = "32bitle"
	} else if opts == "-eS" {
		ext = "8bit"
	} else if opts == "-es" {
		ext = "7bit"
	} else {
		ext = "si"
	}

	res = combine(res, lres)
	_, fname := filepath.Split(filePath)
	ioutil.WriteFile(fmt.Sprintf("%s.str.%s", fname, ext), res, 0777)
}
