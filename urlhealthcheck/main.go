package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

var Revision string

var Urls []string

var wg sync.WaitGroup

func healthChecker(url string) {
	if !strings.HasPrefix(url, "http") {
		url = "https://" + url
	}
	client := &http.Client{
		Timeout: 20 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		log.Printf("The Website  %v is currently OFF  ", url, err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		log.Printf("The Website %v is currently ON", url)
	} else {
		log.Printf("The Website %v is currently OFF (HTTP Status: %d)", url, resp.StatusCode)
	}
	wg.Done()
}

func main() {

	fd, err := os.Open("data.txt")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("successfully Open File")
	defer fd.Close()

	scanner := bufio.NewScanner(fd)

	for scanner.Scan() {
		line := scanner.Text()
		Urls = append(Urls, line)
	}

	wg.Add(len(Urls))

	for _, url := range Urls {
		go healthChecker(url)
	}
	wg.Wait()
}
