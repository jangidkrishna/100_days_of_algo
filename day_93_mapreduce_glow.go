package main

import (
    "strings"
    "github.com/chrislusf/glow/flow"
)

func main() {
    flow.New().TextFile(
        "/home/n3m0/txt.txt", 3,
    ).Filter(func(line string) bool {
        return !strings.HasPrefix(line, "#")
    }).Map(func(line string, ch chan string) {
        for _, token := range strings.Split(line, " ") {
            ch <- token
        }
    }).Map(func(key string) int {
        return 1
    }).Reduce(func(x int, y int) int {
        return x + y
    }).Map(func(x int) {
        println("count:", x)
    }).Run()
}
