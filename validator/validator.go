package validator

import (
	"regexp"
)

func Required(value string) (result bool, errMsg string) {
	result = (len(value) != 0)
	if !result {
		errMsg = "パスワードが入力されていません"
	}
	return
}

func CharType(value string) (result bool, errMsg string) {
	result = (len(value) == 0)
	if result {
		return
	}
	result, _ = regexp.MatchString("^[0-9a-zA-Z!-/:-@¥[-`{-~]+$", value)
	if !result {
		errMsg = "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	}
	return
}

func MaxLength(value string) (result bool, errMsg string) {
	result = (len(value) <= 256)
	if !result {
		errMsg = "パスワードは256文字以内で入力してください"
	}
	return
}
