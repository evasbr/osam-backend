package errors

type HttpError struct {
	StatusCode int
	Messages   []string
}

func (h HttpError) Error() string {
	if len(h.Messages) == 0 {
		return "Unexpected error"
	}

	return h.Messages[0]
}
