package static

import "errors"

var (
	ErrReadTagID = errors.New("tagid read failed")
	ErrHasPosts  = errors.New("tag has posts")
)
