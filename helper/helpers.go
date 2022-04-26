package helper

import (
	"fmt"
	"regexp"
)

// commandError is an error used to signal different error situations in command handling.
type commandError struct {
	s         string
	userError bool
}

func (c commandError) Error() string {
	return c.s
}

func (c commandError) isUserError() bool {
	return c.userError
}

func NewUserError(a ...any) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: true}
}

func NewSystemError(a ...any) commandError {
	return commandError{s: fmt.Sprintln(a...), userError: false}
}

func NewSystemErrorF(format string, a ...any) commandError {
	return commandError{s: fmt.Sprintf(format, a...), userError: false}
}

// Catch some of the obvious user errors from Cobra.
// We don't want to show the usage message for every error.
// The below may be to generic. Time will show.
var userErrorRegexp = regexp.MustCompile("unknown flag")

func isUserError(err error) bool {
	if cErr, ok := err.(commandError); ok && cErr.isUserError() {
		return true
	}

	return userErrorRegexp.MatchString(err.Error())
}
