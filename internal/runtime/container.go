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
