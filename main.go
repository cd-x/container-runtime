package main

import (
	r "container-runtime/internal/runtime"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const OLDROOT = "/oldroot"

func main() {

	if os.Getenv("CONTAINER_INIT") == "1" {
		r.RunContainerInit()
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
		Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWUSER,
		UidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getuid(), Size: 1},
		},
		GidMappings: []syscall.SysProcIDMap{
			{ContainerID: 0, HostID: os.Getgid(), Size: 1},
		},
		GidMappingsEnableSetgroups: false,
	}
	cmd.Env = append(os.Environ(), "CONTAINER_INIT=1")

	log.Println("Parent PID: ", os.Getpid())

	log.Println("starting container...")
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}
