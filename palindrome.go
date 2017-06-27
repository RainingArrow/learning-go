package main

import (
	"log"
	"regexp"
	"strings"
)

var re *regexp.Regexp

func init() {
	var err error

	re, err = regexp.Compile("[^a-z]+")
	if err != nil {
		log.Fatal(err)
	}
}

func isPalindrome(line string) bool {
	line = strings.ToLower(line)
	line = re.ReplaceAllString(line, "")

	mid := len(line) / 2
	last := len(line) - 1
	for i := 0; i < mid; i++ {
		if line[i] != line[last-i] {
			return false
		}
	}
	return true
}
