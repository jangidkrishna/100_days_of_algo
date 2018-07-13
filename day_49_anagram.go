package main

import (
    "fmt"
    "sort"
    "strings"
)

func format(str string) string {
    s := strings.Split(str, "")
    sort.Strings(s)
    return strings.Join(s, "")
}

func anagram (str1 string, str2 string) bool {
    str_1 := format(strings.ToLower(str1))
    str_2 := format(strings.ToLower(str2))
    return str_1 == str_2
}

func main() {
    if (anagram("mode","dome")) {
        fmt.Println("anagrams")
    } else {
        fmt.Println("not anagrams")
    }

}
