package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

const (
	ColorOk    = "\x1b[32m"
	ColorWarn  = "\x1b[33m"
	ColorErr   = "\x1b[31m"
	ColorReset = "\x1b[0m"
)

func checkLink(link string) {
	startTime := time.Now()
	resp, err := http.Get(link)
	elapsed := time.Since(startTime)

	if err != nil {
		fmt.Println(link + "\t" + ColorErr + "✕" + ColorReset + "\t" + "ERROR" + "\t" + elapsed.String() + "\t" + err.Error())
	} else {
		shapeColor := ColorOk
		if resp.StatusCode != http.StatusOK {
			shapeColor = ColorErr
		}
		fmt.Println(link + "\t" + shapeColor + "✔" + ColorReset + "\t" + resp.Status + "\t" + elapsed.String())
	}
}

func main() {
	links := []string{
		"https://noahtigner.com",
		"http://noahtigner.com",
		// "https://noahtigners.com",
	}

	for {
		var waitGroup sync.WaitGroup
		startTime := time.Now()

		// check each link concurrently
		for _, link := range links {
			waitGroup.Go(func() {
				checkLink(link)
			})
		}

		// wait for the batch to complete
		waitGroup.Wait()

		// wait another ~5 seconds before starting again
		secondsToSleep := 5*time.Second - time.Since(startTime)
		time.Sleep(secondsToSleep)
	}
}
