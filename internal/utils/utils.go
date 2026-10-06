package utils

import (
	"os"
	"path/filepath"
)

const ROOT_FS = "/rootfs"

func GetNewRootPath() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(cwd, ROOT_FS)
}
