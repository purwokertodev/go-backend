package model

import "errors"

type Identity struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	PasswordSalt string `json:"passwordSalt"`
}

func (i *Identity) IsValidPassword(password string) error {
	if i.Password != password {
		return errors.New("Password is not valid")
	}
	return nil
}
