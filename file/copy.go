package file

import (
	"io"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

var (
	ErrNotFindPath        = errors.New("path was not found")
	ErrNotDirectory       = errors.New("path was not directory")
	ErrNotExistOptionFile = errors.New("not found file which is option of the minecraft")
	ErrFileOpenFailed     = errors.New("failed to open file")
	ErrFileCreationFailed = errors.New("failed to create file")
	ErrCopyFailed         = errors.New("failed to copy file")
)

const optionFileName = "option.txt"

// CopyOption copies option file to target directory
func CopyOption(source string, target string) error {

	if err := checkValidCopyPaths(source); err != nil {
		return errors.Wrap(err, "provided path is not a valid directory or does not exist")
	}
	if err := checkValidCopyPaths(target); err != nil {
		return errors.Wrap(err, "provided path is not a valid directory or does not exist")
	}

	sourcePath := convertToAbsForOption(source)
	targetPath := convertToAbsForOption(target)

	if err := copyToNewer(sourcePath, targetPath); err != nil {
		return errors.Wrap(err, "failed to copy file")
	}

	return nil
}

func checkValidCopyPaths(path string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) {
		return ErrNotFindPath
	} else if err != nil && !os.IsNotExist(err) {
		return errors.Wrap(err, "unknown error was gotten")
	} else if !f.IsDir() {
		return ErrNotDirectory
	}
	return nil
}

func convertToAbsForOption(path string) string {
	var absPath string
	if abs, err := filepath.Abs(path); err != nil {
		absPath = path
	} else {
		absPath = abs
	}
	return filepath.Join(absPath, optionFileName)
}

func copyToNewer(source string, target string) error {
	var (
		src *os.File
		dst *os.File
	)

	src, err := os.Open(source)
	if err != nil {
		if os.IsNotExist(err) {
			return ErrNotExistOptionFile
		}
		return ErrFileOpenFailed
	}
	defer src.Close()

	dst, err = os.Create(target)
	if err != nil {
		return ErrFileCreationFailed
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return ErrCopyFailed
	}

	return nil
}
