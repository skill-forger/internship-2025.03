package hashing

import (
	"golang.org/x/crypto/bcrypt"
)

// Bcrypt is an implementation of hashing.Algorithm
type Bcrypt struct {
	cost int
}

// NewBcrypt creates and returns a Bcrypt implementation of the hashing Algorithm
func NewBcrypt() Algorithm {
	return &Bcrypt{cost: bcrypt.DefaultCost}
}

// Generate returns the hashed value from the input plain value
func (b *Bcrypt) Generate(value []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(value, b.cost)
}

// Compare checks the hashed value with the plain value from the input
// returns ErrHashingComparisonMismatch error if comparison is mismatched
func (b *Bcrypt) Compare(hashedValue, plainValue []byte) error {
	err := bcrypt.CompareHashAndPassword(hashedValue, plainValue)
	if err != nil {
		return ErrHashingComparisonMismatch
	}

	return nil
}
