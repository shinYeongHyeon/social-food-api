package user_domain_test

import (
	"social-food-api/src/user/domain"
	"testing"
)

func TestUserIdSuccessRegex(t *testing.T) {
	var emails = []string{"test@gmail.com", "test_-1@gmail.com", "test+123@gmail.kr", "pxZYHKy@KxmqmDA.net"}

	for _, email := range emails {
		_, err := user_domain.CreateUserId(email)

		if err != nil {
			t.Errorf("Email validator was not working : %s", email)
		}
	}
}

func TestUserIdFailedRegex(t *testing.T) {
	var badCases = []string{"test#@gmail.com", "badexample", "bad@gmailcom", "bad@gmail.c"}

	for _, email := range badCases {
		_, err := user_domain.CreateUserId(email)

		if err == nil {
			t.Errorf("Email validator was not working : %s", email)
		}
	}
}
