package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var input []string
	for i := 0; i < 10000; i++ {
		input = append(input, fmt.Sprintln(i))
	}
	start1 := time.Now()
	useLoop(input)
	elapse1 := time.Since(start1)
	fmt.Println("Loop: ", elapse1)

	start2 := time.Now()
	useStringsJoin(input)
	elapse2 := time.Since(start2)
	fmt.Println("String Join: ", elapse2)

	faster := float64(elapse1.Nanoseconds()) / float64(elapse2.Nanoseconds())
	fmt.Printf("Strigns Join is faster than Loop %.2f times.", faster)
}

func useLoop(input []string) string {
	var result string = ""
	for _, str := range input {
		result += str + ", "
	}
	return result
}

func useStringsJoin(input []string) string {
	return strings.Join(input, ", ") + ", "
}
