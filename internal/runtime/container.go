package runtime

import (
	"fmt"
	"syscall"
)

func SetupProc(rootfs string) error {
	fmt.Println("mounting /proc...")

	return syscall.Mount(
		"proc",
		rootfs+"/proc",
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
	fmt.Println("pivoting root to:", newRoot)

	// pivot_root requires newRoot to be a mount point (you already bind-mount it)
	if err := syscall.Chdir(newRoot); err != nil {
		return fmt.Errorf("chdir newroot: %w", err)
	}

	// Stack the old root on top of the new root at the same location
	if err := syscall.PivotRoot(".", "."); err != nil {
		return fmt.Errorf("pivot_root: %w", err)
	}

	// "." now refers to the OLD root (it's the top mount). Detach it.
	if err := syscall.Unmount(".", syscall.MNT_DETACH); err != nil {
		return fmt.Errorf("unmount old root: %w", err)
	}

	return syscall.Chdir("/")
}

func MountDirectory(directory string) error {
	fmt.Printf("mounting directory %s ...\n", directory)
	return syscall.Mount(directory, directory, "", syscall.MS_BIND, "")
}

func UnmountOldRoot(directory string) error {
	fmt.Printf("unmounting directory %s ...\n", directory)
	return syscall.Unmount(directory, syscall.MNT_DETACH)
}
