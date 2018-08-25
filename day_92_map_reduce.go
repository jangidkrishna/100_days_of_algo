package main

import (
	"fmt"
	"sync"
)

func mapp(in <-chan string, out chan<- map[string]int) {
	counter := map[string]int{}
	for words := range in {
		counter[words] = counter[words] + 1
	}
	out <- counter
	close(out)
}

func reduce(in <-chan int, out chan<- float32) {
	sum, counter := 0, 0
	for n := range in {
		sum += n
		counter++
	}
	out <- float32(sum) / float32(counter)
	close(out)
}

func input(out [3]chan<- string) {
	inp := [][]string{
		{"noun", "verb", "verb", "noun", "verb", "noun", "verb", "noun"},
		{"noun", "verb", "verb", "verb", "noun", "noun", "verb", "noun", "noun", "verb", "noun"},
		{"noun", "verb", "verb", "verb", "verb", "verb"},
	}

	for i := range out {
		go func(ch chan<- string, word []string) {
			for _, w := range word {
				ch <- w
			}
			close(ch)
		}(out[i], inp[i])
	}
}

func shuffle(in []<-chan map[string]int, out [2]chan<- int) {
	var wg sync.WaitGroup
	wg.Add(len(in))
	for _, ch := range in {

		go func(c <-chan map[string]int) {

			for m := range c {
				noun, ok := m["noun"]
				if ok {
					out[0] <- noun
				}
				verb, ok := m["verb"]
				if ok {
					out[1] <- verb
				}
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out[0])
		close(out[1])
	}()
}

func output(in []<-chan float32) {
	var wg sync.WaitGroup
	wg.Add(len(in))

	name := []string{"noun", "verb"}
	for i := 0; i < len(in); i++ {
		go func(n int, c <-chan float32) {
			for avg := range c {
				fmt.Printf("Avg of %ss is %f\n", name[n], avg)
			}
			wg.Done()
		}(i, in[i])
	}
	wg.Wait()
}

func main() {
	n := 10
	txt1 := make(chan string, n)
	txt2 := make(chan string, n)
	txt3 := make(chan string, n)
	map1 := make(chan map[string]int, n)
	map2 := make(chan map[string]int, n)
	map3 := make(chan map[string]int, n)
	reduce1 := make(chan int, n)
	reduce2 := make(chan int, n)
	avg1 := make(chan float32, n)
	avg2 := make(chan float32, n)

	go input([3]chan<- string{txt1, txt2, txt3})
	go mapp(txt1, map1)
	go mapp(txt2, map2)
	go mapp(txt3, map3)
	go shuffle([]<-chan map[string]int{map1, map2, map3}, [2]chan<- int{reduce1, reduce2})
	go reduce(reduce1, avg1)
	go reduce(reduce2, avg2)

	output([]<-chan float32{avg1, avg2})
}
