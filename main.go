package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: container-runtime <command>")
		os.Exit(1)
	}

	cmd := exec.Command(os.Args[1], os.Args[2:]...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}

	// fmt.Println("starting container...")
	fmt.Println("PID: ", os.Getpid())

	if err := cmd.Start(); err != nil {
		panic(err)
	}

	if err := runtime.SetupProc(); err != nil {
		panic(err)
	}

	fmt.Print("child pid: ", cmd.Process.Pid)

	if err := cmd.Wait(); err != nil {
		fmt.Println("child exited: ", err)
	}
}
