package user

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	if email == "" {
		return "", errors.New("user email must not be empty")
	}
	if !validateEmail(email) {
		return "", errors.New("user email is invalid")
	}
	return Email(email), nil
}

func validateEmail(email string) bool {
	// Emailの正規表現パターン
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// 正規表現のコンパイル
	regexp := regexp.MustCompile(pattern)

	// マッチングを行う
	match := regexp.MatchString(email)

	return match
}
