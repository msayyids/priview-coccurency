package main

import (
	"fmt"
	"sync"
	"time"
)

func soal2[T any](input []T, wg *sync.WaitGroup) {
	defer wg.Done()
	for _, v := range input {
		fmt.Println(v)
	}
}

func soal1[T any](input []T) {
	for _, v := range input {
		fmt.Println(v)
	}
	time.Sleep(3 * time.Second)
}

func soal3[T interface{}](input []T, wg *sync.WaitGroup, c chan any) {
	defer wg.Done()

	for _, v := range input {
		c <- v
	}

}

func main() {

	arr1 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	arr2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	c := make(chan any)

	go soal1(arr1)
	go soal1(arr2)

	wg.Add(4)

	go soal2(arr1, &wg)
	go soal2(arr2, &wg)
	go soal3(arr1, &wg, c)
	go soal3(arr2, &wg, c)

	wg.Wait()
	close(c)

	for result := range c {
		fmt.Println(result)
	}

}
