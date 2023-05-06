package user_domain

import (
	"errors"
	"regexp"
)

type UserId struct {
	value string
}

func CreateUserId(userId string) (UserId, error) {
	if !regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-z]{2,4}$`).MatchString(userId) {
		return UserId{}, errors.New("invalid email")
	}

	return UserId{userId}, nil
}

func (userId UserId) Value() string {
	return userId.value
}
