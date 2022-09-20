package fileutils

import (
	"io"
	"os"

	"github.com/chef/automate/lib/platform/sys"
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

func CalDirSizeInGB(path string) (float64, error) {
	size, err := sys.DirSize(path)
	if err != nil {
		return -1, err
	}
	return float64(size) / (1024 * 1024 * 1024), nil
}

func CheckSpaceAvailability(dir string, minSpace float64) (bool, error) {
	dirSpaceInGB, err := GetFreeSpaceinGB(dir)
	if err != nil {
		return false, err
	}
	if minSpace <= dirSpaceInGB {
		return true, nil
	}
	return false, nil
}

func GetFreeSpaceinGB(dir string) (float64, error) {
	v, err := sys.SpaceAvailForPath(dir)
	if err != nil {
		return -1, err
	}
	return float64(v) / (1024 * 1024), nil
}
