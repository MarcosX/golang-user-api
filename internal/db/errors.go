package db

import "fmt"

type ErrUserNotFound struct {
	Id string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("could not find user %s", e.Id)
}
