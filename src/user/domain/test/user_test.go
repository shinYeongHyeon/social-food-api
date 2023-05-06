package user_domain

import (
	"github.com/go-faker/faker/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/shared/domain"
	"social-food-api/src/shared/enums"
	"social-food-api/src/user/domain"
	"testing"
)

func TestCreateUserSuccess(t *testing.T) {
	email := faker.Email()
	from := "KAKAO"
	userName := faker.Name()

	var userId, _ = user_domain.CreateUserId(email)
	var loginMethods, _ = shared_enums.FromString(from)
	var name, _ = shared_domain.CreateName(userName)

	user, err := user_domain.CreateNewUser(user_domain.NewUserProps{
		UserId: userId,
		Name:   name,
		From:   loginMethods,
	})

	if err != nil ||
		user.GetUserId().Value() != email ||
		user.GetFrom().String() != from ||
		user.GetName().Value() != userName ||
		user.GetCreatedAt() > timestamppb.Now().Seconds ||
		user.GetUpdatedAt() > timestamppb.Now().Seconds ||
		user.GetCreatedAt() != user.GetUpdatedAt() {
		t.Fatalf("create failed or get errors")
	}
}
