package main

import (
	r "container-runtime/internal/runtime"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {

	if os.Getenv("CONTAINER_INIT") == "1" {
		runContainerInit()
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("usage: container-runtime <command>")
		os.Exit(1)
	}

	cmd := exec.Command("/proc/self/exe", os.Args[1:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	cmd.Env = append(os.Environ(), "CONTAINER_INIT=1")

	// fmt.Println("starting container...")
	fmt.Println("PID: ", os.Getpid())

	if err := cmd.Run(); err != nil {
		fmt.Println("child exited: ", err)
		os.Exit(1)
	}
}

func runContainerInit() {
	fmt.Println("inside container init")
	fmt.Println("PID:", os.Getpid())

	fmt.Println("setting up /proc...")
	if err := r.SetupProc(); err != nil {
		panic(err)
	}

	fmt.Println("executing: ", os.Args[1:])
	err := syscall.Exec(
		os.Args[1],
		os.Args[1:],
		os.Environ(),
	)

	if err != nil {
		panic(err)
	}
}

func readlink(path string) string {
	value, err := os.Readlink(path)
	if err != nil {
		return err.Error()
	}
	return value
}
