package validator

func Required(value string) (result bool) {
	return result == (len(value) == 0)
}