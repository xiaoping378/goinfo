package main

import "fmt"
import "runtime/trace"
import "os"

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	} else {
		defer f.Close()
	}
	if err := trace.Start(f); err != nil {
		panic(err)
	}
	defer trace.Stop()

	fmt.Println("Hello")

}
