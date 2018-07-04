package main

import (
	"fmt"
	"strings"
)

var (
	morse_alphabet = map[rune]string{
		'A': ".-",
		'B': "-...",
		'C': "-.-.",
		'D': "-..",
		'E': ".",
		'F': "..-.",
		'G': "--.",
		'H': "....",
		'I': "..",
		'J': ".---",
		'K': "-.-",
		'L': ".-..",
		'M': "--",
		'N': "-.",
		'O': "---",
		'P': ".--.",
		'Q': "--.-",
		'R': ".-.",
		'S': "...",
		'T': "-",
		'U': "..-",
		'V': "...-",
		'W': ".--",
		'X': "-..-",
		'Y': "-.--",
		'Z': "--..",
		'1': ".----",
		'2': "..---",
		'3': "...--",
		'4': "....-",
		'5': ".....",
		'6': "-....",
		'7': "--...",
		'8': "---..",
		'9': "----.",
		'0': "-----",
	}

	inverse_alphabet = make(map[string]rune, len(morse_alphabet))
)

func inverse_alphabet_inti() {
	for k, m := range morse_alphabet {
		inverse_alphabet[m] = k
	}
}

func encode_to_morse(str string) string {
	res := make([]string, 0, len(str))
	str = strings.ToUpper(str)

	for _, c := range str {
		if m, ok := morse_alphabet[c]; ok {
			res = append(res, m)
		}
	}
	return strings.Join(res, "  ")
}

func decode(str string) string {
	res := make([]rune, 0, len(str)/3)
	morse_code := strings.Split(str, "  ")
	for _, m := range morse_code {
		if c, ok := inverse_alphabet[m]; ok {
			res = append(res, c)
		}
	}
	return string(res)
}

func main() {
	fmt.Println(encode_to_morse("Hello"))
}
