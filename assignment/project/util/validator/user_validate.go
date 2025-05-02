package validator

import (
	"golang-project/static"
	"net/mail"
	"regexp"
	"strings"
)

var nameRegexp = regexp.MustCompile(`^[\p{L}][\p{L}\s\-']*$`)

// ValidateEmail validate email field
func ValidateEmail(email string) error {

	if _, err := mail.ParseAddress(email); err != nil {
		return static.ErrInvalidEmail
	}
	return nil
}

// ValidateName validate name field
func ValidateName(n string) error {
	n = strings.TrimSpace(n)
	if !nameRegexp.MatchString(n) {
		return static.ErrInvalidName
	}
	return nil
}
