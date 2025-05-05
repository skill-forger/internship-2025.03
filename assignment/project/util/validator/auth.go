package validator

import (
	"net/mail"
	"regexp"
	"strings"

	ct "golang-project/internal/contract"
	"golang-project/static"
)

var nameRegexp = regexp.MustCompile(`^[\p{L}][\p{L}\s\-']*$`)

func ValidateSignUpRequest(request ct.SignUpRequest) error {
	//Validate email
	if _, err := mail.ParseAddress(request.Email); err != nil {
		return static.ErrInvalidEmail
	}

	//Validate name
	lname := strings.TrimSpace(request.LastName)
	fname := strings.TrimSpace(request.FirstName)
	if !(nameRegexp.MatchString(lname) && nameRegexp.MatchString(fname)) {
		return static.ErrInvalidName
	}

	return nil
}
