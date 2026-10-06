package runtime

import (
	"fmt"
	"syscall"
)

func SetupProc() error {
	fmt.Println("mounting /proc...")

	return syscall.Mount(
		"proc",
		"/proc",
		"proc",
		0,
		"",
	)
}

func MakeRootPrivate() error {
	fmt.Println("making root private ...")
	return syscall.Mount("", "/", "", syscall.MS_PRIVATE|syscall.MS_REC, "")
}
