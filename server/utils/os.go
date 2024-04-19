package utils

import "os"

// RemoveIgnoreNonExistent removes a file, ignoring if it doesn't exist.
//
// There are many cases where we just want to make sure a file is gone, and
// don't really care if it never existed to begin with.
func RemoveIgnoreNonExistent(file string) error {
	err := os.Remove(file)
	if err == nil || os.IsNotExist(err) {
		return nil
	}

	return err
}
