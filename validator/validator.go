package validator

func Required(value string) (result bool) {
	result = (len(value) == 0)
	return
}
