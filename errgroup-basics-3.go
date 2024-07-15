package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
errgroup.WithContext() function creates a new group of goroutines of the errgroup.Group type and a new context.Context,
which can be passed between goroutines and will allow canceling the execution of the task group.
*/
func egrpPattern3() {
	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				log.Printf("goroutine 1 should cancel")
				return nil
			default:
				fmt.Println("task executing")

				if _, err := http.Get("hts://blog.kennycoder.io"); err != nil {
					return err
				}
				time.Sleep(1 * time.Second)
			}
		}
	})
	eg.Go(func() error {
		for i := 0; i < 10; i++ {
			select {
			case <-ctx.Done():
				log.Printf("goroutine 2 should cancel")
				return nil
			default:
				_, err := http.Get("https://google.com")
				if err != nil {
					return err
				}
				time.Sleep(6 * time.Second)
			}
		}
		return nil
	})
	if err := eg.Wait(); err != nil {
		log.Fatalf("get error: %v", err)
	}
}
