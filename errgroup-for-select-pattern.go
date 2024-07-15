package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"golang.org/x/sync/errgroup"
)

// errFailure some custom error.
var errFailure = errors.New("some error")

// change test3 to main() before the execution
func test4() {
	// Create errgroup with context.
	group, qctx := errgroup.WithContext(context.Background())

	// Run first periodic task.
	//executes second
	group.Go(func() error {
		fmt.Println("runs first task")
		firstTask(qctx)
		return nil
	})

	// Run second task.
	//executes first
	group.Go(func() error {
		fmt.Println("runs second task")
		if err := secondTask(); err != nil {
			return err
		}
		return nil
	})

	// Wait for all tasks to complete or the error to appear.
	if err := group.Wait(); err != nil {
		fmt.Printf("errgroup tasks ended up with an error: %v", err)
	}
}

func firstTask(ctx context.Context) {
	var counter int
	for {
		select {
		case <-ctx.Done():
			return
		case <-time.After(500 * time.Millisecond):
			fmt.Println("some task")
			if counter > 10 {
				return
			}
			counter++
		}
	}
}

func secondTask() error {
	time.Sleep(7 * time.Second)
	return errFailure
}
