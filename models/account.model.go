package models

import "reflect"

type Account struct {
	ID      int    `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Balance int32  `json:"balance,omitempty"`
}

func (acc Account) IsEmpty() bool {
	return reflect.DeepEqual(acc, Account{})
}
