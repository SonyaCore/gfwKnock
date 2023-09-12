// Package ulimit provides a function to set the ulimit (maximum number of open file descriptors) on Linux and Darwin (macOS) systems.
package ulimit

import (
	"fmt"
	"runtime"
	"syscall"
)

// SetULimit sets the ulimit (maximum number of open file descriptors) to a specified value.
func SetULimit() error {
	if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
		var rLimit syscall.Rlimit
		rLimit.Max = 127000
		rLimit.Cur = 127000

		err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
		if err != nil {
			fmt.Println("Error Setting Rlimit ", err)
			return err
		}

		// Verify that the RLIMIT_NOFILE has been set correctly
		var currentRlimit syscall.Rlimit
		syscall.Getrlimit(syscall.RLIMIT_NOFILE, &currentRlimit)
		fmt.Printf("New max number of open files: %d\n", currentRlimit.Max)
	}
	return nil
}
