package hashing

import "errors"

var (
	ErrHashingComparisonMismatch = errors.New("hashed value and plain value mismatch")
)

// Algorithm represents the hashing asymmetrical encryption algorithm.
type Algorithm interface {
	Generate(value []byte) ([]byte, error)
	Compare(hashedValue, plainValue []byte) error
}
