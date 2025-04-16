package model

// User represents user table from the database
type User struct {
	BaseModel
	FirstName    string
	LastName     string
	Email        string
	Password     string
	Pseudonym    string
	ProfileImage string
	Biography    string
}
