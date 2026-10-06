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

func PivotRoot(newRoot string) error {
	fmt.Println("pivoting root to: ", newRoot)
	return syscall.PivotRoot(newRoot, newRoot+"/oldroot")
}

func MountDirectory(directory string) error {
	fmt.Printf("mounting directory %s ...\n", directory)
	return syscall.Mount(directory, directory, "", syscall.MS_BIND, "")
}
