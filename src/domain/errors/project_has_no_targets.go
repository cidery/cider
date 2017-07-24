package errors

import (
	"errors"
	"fmt"
)

type ProjectHasNoTargetsError struct {
	error
}

func NewProjectHasNoTargetsError() *ProjectHasNoTargetsError {
	return &ProjectHasNoTargetsError{
		errors.New(fmt.Sprintf(
			"Submitted project has no targets",
		)),
	}
}
