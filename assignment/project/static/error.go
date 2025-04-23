package static

import "errors"

var (
	ErrNotFound = errors.New("tag not found")
	ErrHasPosts = errors.New("tag has posts")
)
