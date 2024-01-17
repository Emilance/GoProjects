package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	lines := flag.Bool("l", false, "count number of lines")
	flag.Parse()
	// word := flag.string
	result := count(os.Stdin, *lines)
	fmt.Println(result)

	// var intNum string = " Hllo"  + " mee"
	// fmt.Println(utf8.RuneCountInString(intNum))
}

func count(r io.Reader, lineCount bool) int {

	//read r
	scanner := bufio.NewScanner(r)
	if !lineCount {
		scanner.Split(bufio.ScanWords)
	}
	counter := 0

	for scanner.Scan() {
		counter += 1
	}

	err := scanner.Err()
	if err != nil {
		return 0
	}
	return counter
}
