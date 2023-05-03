package shared_core_test

import (
	"reflect"
	"social-food-api/src/shared/core"
	"testing"
)

func TestCreateUuid(t *testing.T) {
	createdUuid := shared_core.CreateUuid()
	if createdUuid == "" || reflect.TypeOf(createdUuid).String() != "string" {
		t.Fatal("Fail to create uuid")
	}
}
