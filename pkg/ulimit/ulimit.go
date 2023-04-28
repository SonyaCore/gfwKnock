package ulimit

import (
	"fmt"
	"runtime"
	"syscall"
)

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
