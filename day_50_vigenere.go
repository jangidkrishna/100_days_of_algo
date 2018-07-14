package main

import (
	"fmt"
)

func encrypt(str string, key string) string {
	if err := check(key); err != nil {
		fmt.Println("error: ", err.Error())
		return "err"
	}
	j := 0
	n := len(str)
	res := ""
	for i := 0; i < n; i++ {
		r := str[i]
		if r >= 'a' && r <= 'z' {
			res += string((r+key[j]-2*'a')%26 + 'a')
		} else if r >= 'A' && r <= 'Z' {
			res += string((r+key[j]-2*'A')%26 + 'A')
		} else {
			res += string(r)
		}

		j = (j + 1) % len(key)
	}
	return res
}

func decrypt(str string, key string) string {
	if err := check(key); err != nil {
		fmt.Println("Error: ", err.Error())
		return "err"
	}
	res := ""
	j := 0
	n := len(str)
	for i := 0; i < n; i++ {
		r := str[i]
		if r >= 'a' && r <= 'z' {
			res += string((r-key[j]+26)%26 + 'a')
		} else if r >= 'A' && r <= 'Z' {
			res += string((r-key[j]+26)%26 + 'A')
		} else {
			res += string(r)
		}

		j = (j + 1) % len(key)
	}
	return res
}

func check(key string) error {
	if len(key) == 0 {
		return fmt.Errorf("invalid key")
	}
	for _, r := range key {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z') {
			return fmt.Errorf("invalid key")
		}
	}
	return nil
}

func main() {
	var (
		msg, key string
		mode     rune
	)
	fmt.Println("enter str to encrypt | decrypt")
	fmt.Scanf("%s\n", &msg)
	fmt.Println("enter key")
	fmt.Scanf("%s\n", &key)
	fmt.Println("'e' to encrypt | 'd' to decrypt")
	fmt.Scanf("%c", &mode)
	if mode == rune('e') {
		fmt.Println("encrypted msg: ", encrypt(msg, key))
	} else if mode == rune('d') {
		fmt.Println("deccrypted msg: ", decrypt(msg, key))
	} else {
		fmt.Println("Unknown")
	}
}
