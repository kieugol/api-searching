package models

import "reflect"

type User struct {
	ID         int        `json:"_,omitempty"`
	Name       string     `json:"name,omitempty"`
	AccountIDS []string   `json:"_,omitempty"`
	Accounts   []*Account `json:"accounts"`
}

func (user *User) IsEmpty() bool {
	return reflect.DeepEqual(*user, User{}) || user == nil
}
