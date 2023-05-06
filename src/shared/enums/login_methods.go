package shared_enums

import "errors"

type LoginMethods struct {
	from string
}

var (
	Kakao = LoginMethods{"KAKAO"}
	Apple = LoginMethods{"APPLE"}
)

func (loginMethods LoginMethods) String() string {
	return loginMethods.from
}

func FromString(from string) (LoginMethods, error) {
	switch from {
	case Kakao.from:
		return Kakao, nil
	case Apple.from:
		return Apple, nil
	default:
		return LoginMethods{}, errors.New("unknown loginMethods")
	}
}
