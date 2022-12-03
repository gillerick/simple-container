package main

import (
	"fmt"
	"os"
)

// docker run <container> cmd args
// go run main.go run cmd args
func main() {
	// Here, we check if the first command argument is 'run' and invoke the method `run`, otherwise, the program panics
	switch os.Args[1] {
	case "run":
		run()
	default:
		panic("what happened?")
	}

}

func run() {
	// We print out all arguments starting from the second one
	fmt.Printf("running %v\n", os.Args[2:])
}
