package shared_enums_test

import (
	"social-food-api/src/shared/enums"
	"testing"
)

func TestUnknownMethod(t *testing.T) {
	_, err := shared_enums.FromString("unknown")
	_, errLowerCase := shared_enums.FromString("KAKAo")

	if err == nil || errLowerCase == nil {
		t.Errorf("unknown methods throw errors")
	}
}

func TestKnownMethod(t *testing.T) {
	_, kakaoErr := shared_enums.FromString("KAKAO")
	_, appleErr := shared_enums.FromString("APPLE")

	if kakaoErr != nil || appleErr != nil {
		t.Errorf("known methods NOT throw errors")
	}
}
