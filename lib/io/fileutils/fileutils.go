package fileutils

import (
	"fmt"
	"io"
	"os"

	"github.com/chef/automate/lib/platform/sys"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// LogCLose closes the given io.Closer, logging any error.
func LogClose(c io.Closer, log logrus.FieldLogger, msg string) {
	if err := c.Close(); err != nil {
		log.WithError(err).Error(msg)
	}
}

// PathExists returns true if the path exists and false if it doesn't
// exist. An error is returned if an unexpected error occurs.  Callers
// who want behavior similar to Ruby's File.exist? or Rusts'
// path::exists functions can ignore the error.
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false, nil
	}

	if err != nil {
		return false, err
	}

	return true, nil
}

// IsSymlink returns true if the given path is a symbolic link and
// false otherwise.  An error is returned on
func IsSymlink(path string) (bool, error) {
	fileInfo, err := os.Lstat(path)
	if err != nil {
		return false, err
	}

	return (fileInfo.Mode() & os.ModeSymlink) == os.ModeSymlink, nil
}

func CalDiskSizeAndDirSize(data, oldData string) (bool, error) {
	v, err := sys.SpaceAvailForPath(data)
	if err != nil {
		return false, err
	}
	diskSpaceInMb := v / 1024

	size, err := sys.DirSize(oldData)
	if err != nil {
		return false, err
	}

	dirSizeInMb := size / (1024 * 1024)

	msg := fmt.Sprintf("Insufficient disk space for migration.\n%s: %5d MB\n%s: %5d MB\n%s",
		"Space Required", (dirSizeInMb + (dirSizeInMb / 10)),
		"Space Available", diskSpaceInMb,
		"To continue with less memory Please use --skip-storage-check")
	// NewData > olddata + 10%of oldData
	if diskSpaceInMb > (uint64(dirSizeInMb) + uint64(dirSizeInMb/10)) {
		return true, nil
	} else {
		return false, errors.New(msg)
	}
}
