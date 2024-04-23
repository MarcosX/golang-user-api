package db

import "fmt"

type ErrUserNotFound struct {
	Id string
}

func (e *ErrUserNotFound) Error() string {
	return fmt.Sprintf("could not find user %s", e.Id)
}

type ErrUserAlreadyExists struct {
	Id string
}

func (e *ErrUserAlreadyExists) Error() string {
	return fmt.Sprintf("dupulicated user %s", e.Id)
}

type ErrValidationFailed struct {
	Field string
}

func (e *ErrValidationFailed) Error() string {
	return fmt.Sprintf("validation failed for field %s", e.Field)
}
