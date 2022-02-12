package models

// User struct to describe object.
type User struct {
	// ID int `validate:"required"`

	Username string `validate:"required,lte=255"`
	Email    string `validate:"required,lte=255"`
	Password string `validate:"required,lte=255"`
}
