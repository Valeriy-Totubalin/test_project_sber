package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/Valeriy-Totubalin/test_project_sber/pkg/rate_limiter"
)

var wg sync.WaitGroup

func main() {
	limiter := rate_limiter.NewTokenBucket(2, time.Minute)
	worker := rate_limiter.NewWorker(limiter, 2)
	c := make(chan []func())
	worker.DoWork(c)

	callbacks := []func(){
		callback1,
		callback2,
		callback3,
	}

	c <- callbacks
	close(c)

	wg.Add(len(callbacks))
	wg.Wait()
}

func callback1() {
	fmt.Println("called 1")
	wg.Done()
}

func callback2() {
	fmt.Println("called 2")
	wg.Done()
}

func callback3() {
	fmt.Println("called 3")
	wg.Done()
}
