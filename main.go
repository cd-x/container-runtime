package main

import (
	r "container-runtime/internal/runtime"
	u "container-runtime/internal/utils"
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"
)

const OLDROOT = "/oldroot"

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

func runContainerInit() {
	log.Printf("inside container init with PID:%d\n", os.Getpid())

	if err := r.MakeRootPrivate(); err != nil {
		panic(err)
	}

	rootfs := u.GetNewRootPath()

	if err := r.MountDirectory(rootfs); err != nil {
		panic(err)
	}
	fmt.Println("[OK] rootfs mounted")

	if err := r.SetupProc(rootfs); err != nil {
		panic(err)
	}
	log.Println("/proc mounted [OK]")

	if err := r.PivotRoot(rootfs); err != nil {
		panic(err)
	}

	fmt.Println("executing: ", os.Args[1:])
	execErr := syscall.Exec(
		os.Args[1],
		os.Args[1:],
		os.Environ(),
	)

	if execErr != nil {
		panic(execErr)
	}
}

func readlink(path string) string {
	value, err := os.Readlink(path)
	if err != nil {
		return err.Error()
	}
	return value
}
