package group_domain_test

import (
	"github.com/go-faker/faker/v4"
	"google.golang.org/protobuf/types/known/timestamppb"
	"social-food-api/src/group/group_domain"
	"social-food-api/src/shared/domain"
	"testing"
)

func TestCreateGroupSuccess(t *testing.T) {
	userName := faker.Name()
	ownerUuid := "test-uuid"

	var name, _ = shared_domain.CreateName(userName)

	group, err := group_domain.CreateNewGroup(group_domain.NewGroupProps{
		OwnerUuid: ownerUuid,
		Name:      name,
	})

	if err != nil ||
		group.GetOwnerUuid() != ownerUuid ||
		group.GetName().Value() != userName ||
		group.GetCreatedAt() > timestamppb.Now().Seconds ||
		group.GetUpdatedAt() > timestamppb.Now().Seconds ||
		group.GetCreatedAt() != group.GetUpdatedAt() {
		t.Fatalf("create failed or get errors")
	}
}
