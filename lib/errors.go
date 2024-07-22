package lib

import "fmt"

type RuntimeError struct {
	token Token
	Err   error
}

func (e RuntimeError) Error() string {
	return fmt.Sprintf("RuntimeError: %v", e.Err)
}
