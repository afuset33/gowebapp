package checker

import (
	"regexp"
)

func LengthCheck(value string, length int) (result bool) {
	result = len(value) >= length
	return
}

func ComboUpperLowerCase(value string) (result bool) {
	regexchk := regexp.MustCompile("^(?=[a-z])(?=[A-Z])[a-zA-Z]+$")
	result = regexchk.MatchString(value)
	return
}
