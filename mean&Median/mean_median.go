package main

import (
	"log"
	"sort"
	"sync"
	"time"
)

func num(v int, wg *sync.WaitGroup, ch1 chan int) {
	time.Sleep(1 * time.Second)
	ch1 <- v
	wg.Done()
}

func calculation(ch1 chan int, ch2 chan string) {
	arr := make([]int, 0)
	var sum, median, n int
	for v := range ch1 {
		arr = append(arr, v)
		time.Sleep(1 * time.Second)
		log.Println("Input num:", v)
		sort.Ints(arr)
		n = len(arr)
		sum += v
		mean := float32(sum) / float32(n)
		log.Printf("Mean:%.2f", mean)
		if n%2 == 0 {
			mid := n / 2
			median = (arr[mid] + arr[mid-1]) / 2
			log.Println("Median:", median)
		} else {
			mid := (n - 1) / 2
			median = arr[mid]
			log.Println("Median:", median)
		}
	}
	ch2 <- "Execution Completed"
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan string)
	go calculation(ch1, ch2)
	var wg sync.WaitGroup
	for i := 1; i < 15; i++ {
		wg.Add(1)
		go num(i, &wg, ch1)
	}
	wg.Wait()
	close(ch1)
	log.Println(<-ch2)
}
