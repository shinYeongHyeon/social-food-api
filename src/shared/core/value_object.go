package shared_core

// ValidateValueObjectErrors : ValueObject 생성시 오류 났을 경우 Response
func ValidateValueObjectErrors(voErrors []error) (bool, error) {
	var occurredError error

	for _, err := range voErrors {
		if err != nil {
			occurredError = err
			break
		}
	}

	return occurredError != nil, occurredError
}
