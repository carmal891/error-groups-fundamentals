package main

import (
	"context"
	"fmt"
	"net/http"
	"runtime"

	"golang.org/x/sync/errgroup"
)

/*
The main idea of the errgroup pattern is to start a group of goroutines,
wait for them to finish their work, and handle any errors that may occur during execution. */

// change test to main() before the execution
func test() {

	urls := []string{
		"https://www.easyjet.com/",
		"https://ww.sskyscanner.de/",
		"https://ww.swiss.com/",
		"https://www.ryanair.com",
		"https://wizzair.com/",
	}

	ctx := context.Background()
	g, _ := errgroup.WithContext(ctx)

	// See https://pkg.go.dev/golang.org/x/sync/errgroup#Group.SetLimit
	g.SetLimit(2)

	for _, url := range urls {
		url := url // https://golang.org/doc/faq#closures_and_goroutines
		g.Go(func() error {
			fmt.Printf("%s: checking\n", url)
			res, err := http.Get(url)
			if err != nil {
				return err
			}

			defer res.Body.Close()

			return nil
		})
	}

	fmt.Println("Number of goroutines", runtime.NumGoroutine())

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v", err)
		return
	}

	fmt.Println("Successfully fetched all URLs.")
}
