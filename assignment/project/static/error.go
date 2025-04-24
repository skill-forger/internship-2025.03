package static

import "errors"

var (
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("tag has posts")
	ErrTagNotFound = errors.New("tag not found")
)
