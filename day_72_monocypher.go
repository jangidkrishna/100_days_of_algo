package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

var p = [26]rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i',
	'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}

var ch = [26]rune{'Q', 'W', 'E', 'R', 'T', 'Y', 'U', 'I', 'O',
	'P', 'A', 'S', 'D', 'F', 'G', 'H', 'J', 'K', 'L', 'Z', 'X', 'C', 'V', 'B', 'N', 'M'}

func encrypt(s string) string {
	var c string
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	safe := reg.ReplaceAllString(s, "")

	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if strings.ToLower(string(safe[i])) == string(p[j]) {
				c = c + string(ch[j])
			}
		}
	}
	return c
}

func dec(s string) string {
	var c string
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(p); j++ {
			if string(ch[j]) == string(s[i]) {
				c = c + string(p[j])
			}
		}
	}
	return c
}

func main() {
	s := encrypt("str")
	fmt.Println(s)
	s = dec(s)
	fmt.Println(s)
}
