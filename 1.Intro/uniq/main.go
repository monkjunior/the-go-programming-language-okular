package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	// version1(os.Stdin)
	// version2(os.Args[1:], counts)
	version3(os.Args[1:], counts)

	for line, count := range counts {
		fmt.Printf("%s: %d\n", line, count)
	}
}

func version3(filePath []string, counts map[string]int) {
	for _, file := range filePath {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
}

func version2(filePath []string, counts map[string]int) {

	for _, file := range filePath {
		f, err := os.Open(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%+v", err)
			continue
		}
		myscanner := bufio.NewScanner(f)
		for myscanner.Scan() {
			counts[myscanner.Text()]++
		}
		f.Close()
	}

	//Ignore potential error from myscanner.Err()

}

func version1(reader io.Reader) {
	counts := make(map[string]int)
	myscanner := bufio.NewScanner(reader)
	for myscanner.Scan() {
		if myscanner.Text() == "" {
			break
		}
		counts[myscanner.Text()]++
	}
	//Ignore potential error from myscanner.Err()
}
