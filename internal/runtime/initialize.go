package runtime

import (
	u "container-runtime/internal/utils"
	"fmt"
	"log"
	"os"
	"syscall"
)

func RunContainerInit() {
	log.Printf("inside container init with PID:%d\n", os.Getpid())

	if err := MakeRootPrivate(); err != nil {
		panic(err)
	}

	rootfs := u.GetNewRootPath()

	if err := MountDirectory(rootfs); err != nil {
		panic(err)
	}
	fmt.Println("[OK] rootfs mounted")

	if err := SetupProc(rootfs); err != nil {
		panic(err)
	}
	log.Println("/proc mounted [OK]")

	if err := PivotRoot(rootfs); err != nil {
		panic(err)
	}

	if err := syscall.Sethostname([]byte("container")); err != nil {
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
