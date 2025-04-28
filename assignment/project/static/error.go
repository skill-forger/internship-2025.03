package static

import "errors"

var (
	//Tag error
	ErrReadTagID   = errors.New("error get tag detail")
	ErrHasPosts    = errors.New("error delete tag because it has associated posts")
	ErrTagNotFound = errors.New("error tag id not found")

	//post error
	ErrTagNotFoundOrDeleted = errors.New("error tag id not found or deleted")
	ErrInsertPost           = errors.New("error insert post")
)
