package static

// PaginationDefault defines a struct that holds default pagination values.
type PaginationDefault struct {
	DefaultPage     int
	DefaultPageSize int
}

// Pagination represents the default pagination settings
var Pagination = PaginationDefault{
	DefaultPage:     1,
	DefaultPageSize: 10,
}
