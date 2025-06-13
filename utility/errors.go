package utility

import (
	"fmt"
	"os"
)

func Must[T any](val T, err error) T {
	if err != nil {
		fmt.Printf("%v\n", err)

		os.Exit(1)
	}

	return val
}
