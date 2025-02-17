package error

func ErrMapping(err error) bool {
	allErrors := make([]error, 0)
	allErrors = append(append(GeneralErrors[:], UserErrors[:]...))

	for _, e := range allErrors {
		if err.Error() == e.Error() {
			return true
		}
	}
	return false
}
