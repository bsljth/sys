package sys

import "fmt"

func Log(args ...any) (n int, err error) {
	return fmt.Println(args...)
}