package main

import (
	"fmt"
	"sync"
	"time"
)

func letters(input []string) {
	for _, v := range input {
		fmt.Println(v)
	}
}

func numbers(input []int) {
	for _, v := range input {
		fmt.Println(v)
	}
}

func main() {
	letter := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	number := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup

	go func() {
		letters(letter)
		numbers(number)
	}()

	time.Sleep(2 * time.Second)

	wg.Add(1)

	go func() {
		letters(letter)
		numbers(number)
		wg.Done()
	}()

	wg.Wait()

}
