package stringsly

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"unicode/utf8"
)

var (
	currentOS string
)

func combine(b1 []byte, b2 []byte) []byte {

	newslice := make([]byte, len(b1)+len(b2))
	copy(newslice, b1)
	copy(newslice[len(b1):], b2)
	return newslice
}

func ParseStrings(filePath string) {
	rfile, err := os.Open(filePath)
	defer rfile.Close()

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(rfile)

	for scanner.Scan() {
		b := scanner.Bytes()

		for len(b) > 0 {
			r, size := utf8.DecodeRune(b)
			fmt.Printf("%c %v\n", r, size)

			b = b[size:]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Println(err)
	}
}

func RunStrings(filePath string, incLoc bool, opts string) {
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
		ext = "16bitbe"
	} else if opts == "-eS" {
		ext = "8bit"
	} else if opts == "-es" {
		ext = "7bit"
	} else {
		ext = "si"
	}

	res = combine(res, lres)

	if len(res) > 0 {
		_, fname := filepath.Split(filePath)
		ioutil.WriteFile(fmt.Sprintf("%s.%s", fname, ext), res, 0777)
	} else {
		fmt.Printf("No entries for %s\n", ext)
	}
}
