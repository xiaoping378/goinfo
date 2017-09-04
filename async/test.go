package main

import (
	"fmt"
	"os"
	"time"

	"github.com/xiaoping378/goinfo/async/runner"
)

func main() {
	fmt.Println("Starting...")
	r := runner.New(3 * time.Second)
	r.Add(createTask(), createTask())

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrInterrupt:
			fmt.Println("接收到了中断")
			os.Exit(1)
		case runner.ErrTimeOut:
			fmt.Println("超时了")
			os.Exit(2)

		}
	}
	fmt.Println("Ending...")
}

func createTask() func(int) {
	return func(id int) {
		time.Sleep(4 * time.Second)
		fmt.Println("i'm id ", id)
	}
}
