package models

import "autumn/pkg/autumncore/generic"

type User struct {
	ID       uint
	Username string
	Name     string
	Email    string
	NIK      string
	*generic.ModelImpl[User]
}

func NewUser() *User {
	return &User{}
}
