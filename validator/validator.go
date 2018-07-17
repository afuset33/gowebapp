package validator

import (
	"regexp"
)

func Required(value string) (ok bool, errMsg string) {
	ok = (len(value) != 0)
	if !ok {
		errMsg = "パスワードが入力されていません"
	}
	return
}

func CharType(value string) (ok bool, errMsg string) {
	ok = (len(value) == 0)
	if ok {
		return
	}
	ok, _ = regexp.MatchString("^[0-9a-zA-Z!-/:-@¥[-`{-~]+$", value)
	if !ok {
		errMsg = "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	}
	return
}

func MaxLength(value string) (ok bool, errMsg string) {
	ok = (len(value) <= 256)
	if !ok {
		errMsg = "パスワードは256文字以内で入力してください"
	}
	return
}
