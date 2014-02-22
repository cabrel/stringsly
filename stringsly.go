package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
)

var res []byte

func runStrings(filePath string, opts string) {
	argv := []string{
		"-a",
		"-td",
		opts,
		filePath,
	}

	cmd := exec.Command("strings", argv...)
	lres, err := cmd.CombinedOutput()

	if err != nil {
		log.Fatal(err)
	}

	var header []byte

	if opts == "-el" {
		header = []byte("------ 16-bit (strings -el) ------\n")
	} else if opts == "-eL" {
		header = []byte("------ 32-bit (strings -eL) ------\n")
	} else if opts == "-eB" {
		header = []byte("------ 32-bit (strings -eB) ------\n")
	} else if opts == "-eb" {
		header = []byte("------ 16-bit (strings -eb) ------\n")
	} else if opts == "-eS" {
		header = []byte("------ 8-bit (strings -eS) ------\n")
	} else if opts == "-es" {
		header = []byte("------ 7-bit (strings -es) ------\n")
	}

	res = combine(res, combine(header, lres))
}

func main() {
	flag.Parse()
	fileToString := flag.Arg(0)

	runStrings(fileToString, "-eL")
	runStrings(fileToString, "-el")
	runStrings(fileToString, "-eB")
	runStrings(fileToString, "-eb")
	runStrings(fileToString, "-eS")
	runStrings(fileToString, "-es")

	_, fname := filepath.Split(fileToString)
	ioutil.WriteFile(fmt.Sprintf("%s.str", fname), res, 0777)
}

func combine(b1 []byte, b2 []byte) []byte {

	newslice := make([]byte, len(b1)+len(b2))
	copy(newslice, b1)
	copy(newslice[len(b1):], b2)
	return newslice
}
