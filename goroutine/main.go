package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	start := time.Now()
	userName := fetchUser()

	wg := &sync.WaitGroup{}

	wg.Add(2)
	rspchn := make(chan any, 2)
	go fetchUserLikes(userName, rspchn, wg)
	go fetchUserMatch(userName, rspchn, wg)
	wg.Wait()
	close(rspchn)
	for resp := range rspchn {
		fmt.Println("resp: ", resp)
	}

	// fmt.Println("Likes: ", likes)
	// fmt.Println("match: ", match)
	fmt.Println("Took: ", time.Since(start))

}

func fetchUser() string {
	time.Sleep(time.Millisecond * 100)

	return "BOB"
}

func fetchUserLikes(userName string, rspchn chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 150)

	rspchn <- 11
	wg.Done()
}

func fetchUserMatch(userName string, rspchn chan any, wg *sync.WaitGroup) {
	time.Sleep(time.Millisecond * 100)

	rspchn <- "ANNA"
	wg.Done()
}
