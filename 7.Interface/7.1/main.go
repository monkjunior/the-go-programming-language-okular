package main

import (
	"bufio"
	"fmt"
	"strings"
)

type ByteCounter int
func (c *ByteCounter) Write(p []byte) (int, error){
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int
func (w *WordCounter) Write(p []byte) (int, error){
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		*w ++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return len(p), nil
}

type LineCounter int
func (l *LineCounter) Write(p []byte) (int, error){
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		*l ++
	}

	if err := scanner.Err(); err != nil {
		return -1, err
	}
	return len(p), nil
}


func main() {
	var name string = "VuSon"
	var counter ByteCounter = 0
	fmt.Fprintf(&counter, "My name is %s", name)
	fmt.Println("Bytes", counter)

	var words WordCounter = 0
	fmt.Fprintf(&words, "My real name is %s", name)
	fmt.Println("Words", words)

	var lines LineCounter =0
	var paragraph string =
		"This is line 1\n" +
		"This is line 2\n" +
		"This is line 3"
	fmt.Fprintf(&lines, paragraph)
	fmt.Println("Lines", lines)
	fmt.Println(paragraph)
}