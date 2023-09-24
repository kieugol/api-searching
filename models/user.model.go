package models

import "reflect"

type User struct {
	Name     string     `json:"name,omitempty"`
	Accounts []*Account `json:"accounts,omitempty"`
}

func (user User) IsEmpty() bool {
	return reflect.DeepEqual(user, User{})
}
