package pagination

// CalculateOffset calculates the starting offset for a paginated query.
func CalculateOffset(page int, pageSize int) int {
	return (page - 1) * pageSize
}
