package booking

import "github.com/google/uuid"

type UserType int

const (
	Owner UserType = iota
	Customer
)

const (
	OpNotSupported = "Operation not supported with the current role"
)

var userId = 0

type User struct {
	userId   string
	userType UserType
}

func NewUser(userType UserType) *User {
	return &User{
		userId:   uuid.NewString(),
		userType: userType,
	}
}
