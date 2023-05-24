package user

import "errors"

type Name string

func NewName(name string) (Name, error) {
	if name == "" {
		return "", errors.New("user name must not be empty")
	}
	return Name(name), nil
}
