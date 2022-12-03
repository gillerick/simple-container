//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

// docker run <container> cmd args
// go run main.go run cmd args
func main() {
	// Here, we check if the first command argument is 'run' and invoke the method `run`, otherwise, the program panics
	switch os.Args[1] {
	case "run":
		run()
	case "child":
		child()
	default:
		panic("what happened?")
	}

}

func run() {
	// We print out all arguments starting from the second one
	fmt.Printf("running %v as PID %v\n", os.Args[2:], os.Getpid())

	// Here, we are running command 2, and optionally 3 onwards
	cmd := exec.Command("proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.SysProcAttr = &syscall.SysProcAttr{
		CloneFlags: syscall.RTF_CLONING,
	}
	// We are running the command we set up on line 28
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed with error %v", err)
		return
	}

}

func child() {
	// We print out all arguments starting from the second one
	fmt.Printf("running %v as PID %v\n", os.Args[2:], os.Getpid())

	// Here, we are running command 2, and optionally 3 onwards
	cmd := exec.Command(os.Args[2], os.Args[3:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		log.Fatalf("Command failed with error %v", err)
		return
	}

}
