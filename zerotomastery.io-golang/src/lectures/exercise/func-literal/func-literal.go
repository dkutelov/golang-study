//--Summary:
//  Create a program that can create a report of rune information from
//  lines of text.
//
//--Requirements:
//* Create a single function to iterate over each line of text that is
//  provided in main().
//  - The function must return nothing and must execute a closure
//* Using closures, determine the following information about the text and
//  print a report to the terminal:
//  - Number of letters
//  - Number of digits
//  - Number of spaces
//  - Number of punctuation marks
//
//--Notes:
//* The `unicode` stdlib package provides functionality for rune classification

package main

import (
	"fmt"
	"unicode"
)

type LineCallback func(line string)

func lineIterator(lines []string, callback LineCallback) {
	for _, line := range lines {
		callback(line)
	}
}

func main() {
	lines := []string{
		"There are",
		"68 letters,",
		"five digits,",
		"12 spaces,",
		"and 4 punctuation marks in these lines of text!",
	}

	letters := 0
	digits := 0
	spaces := 0
	punctuation := 0

	lineFunc := func(text string) {
		for _, r := range text {
			if unicode.IsLetter(r) {
				letters += 1
			} else if unicode.IsDigit(r) {
				digits += 1
			} else if unicode.IsSpace(r) {
				spaces += 1
			} else if unicode.IsPunct(r) {
				punctuation += 1
			}
		}
	}

	lineIterator(lines, lineFunc)

	fmt.Printf("The text contains: %v letter, %v digits, %v spaces, %v punctuations.\n", 
	letters, 
	digits, 
	spaces,
	punctuation)
}
