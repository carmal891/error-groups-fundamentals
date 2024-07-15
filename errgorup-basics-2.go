package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

type validationError string

func (v validationError) Error() string {
	return string(v)
}

const validationError1 = validationError("Some validation error 1")
const validationError3 = validationError("Some validation error 3")

/*
use group.Wait() to wait for all tasks in the group to complete.
If any of the tasks finish with an error, the group.Wait() method will return the first error it received.
If all tasks complete successfully, the group.Wait() method will return nil.
*/
func egrpPattern2() {

	egroup := errgroup.Group{}

	egroup.Go(func() error {
		fmt.Println("perform some task 1")
		return validationError1
	})

	egroup.Go(func() error {
		fmt.Println("perform some task 2")
		return nil
	})

	egroup.Go(func() error {
		fmt.Println("perform some task 3")
		return validationError3
	})

	if err := egroup.Wait(); err != nil {
		fmt.Println(err.Error())
	}

}
