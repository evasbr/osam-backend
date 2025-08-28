package errors

import "strings"

type ValidationError struct {
	Messages []string
}

func (v ValidationError) Error() string {
	return strings.Join(v.Messages, ";")
}
