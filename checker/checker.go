package checker

import (
	"log"
	"regexp"
	"strconv"
)

func LengthCheck(value string, length int) (result bool) {
	result = len(value) >= length
	log.Printf("LengthCheck: result-%s", strconv.FormatBool(result))
	return
}

/*
ComboUpperLowerCase
対象文字列が大文字英字と小文字英字を含んでいるかを調べます
TODO:肯定先読みができるライブラリを使う（^(?=.*[a-z]+)(?=.*[A-Z]+)[a-zA-Z]+$）
*/
func ComboUpperLowerCase(value string) (result bool) {
	isUpper, _ := regexp.MatchString(".*[A-Z]+", value)
	isLower, _ := regexp.MatchString(".*[a-z]+", value)
	result = isUpper && isLower
	log.Printf("ComboUpperLowerCase: result-%s", strconv.FormatBool(result))
	return
}

/*
ComboCharaType
対象文字列が半角英字、半角数字、半角記号の組み合わせを含んでいるかを調べます
引数で指定した数以上の文字種を含んでいることが、判定結果OKとなるための条件です。
*/
func ComboCharaType(value string, condition int) (result bool) {
	includeHalfWidthAlpha, _ := regexp.MatchString(".*[a-zA-Z]+", value)
	includeHalfWidthDigit, _ := regexp.MatchString(".*[0-9]+", value)
	includeHalfWidthSign, _ := regexp.MatchString(".*[!-/:-@¥[-`{-~]+", value)
	satisfy := 0
	if includeHalfWidthAlpha {
		satisfy++
	}
	if includeHalfWidthDigit {
		satisfy++
	}
	if includeHalfWidthSign {
		satisfy++
	}
	result = satisfy >= condition
	log.Printf("ComboCharaType: result-%s satisfy-%s", strconv.FormatBool(result), strconv.Itoa(satisfy))
	return
}

/*
ContinuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
*/
func ContinuousChar(value string, condition int) (result bool) {
	regexStr := ".*(.)"
	for i := 0; i < 3; i++ {
		regexStr += "\\1"
	}
	regexStr += ".*"
	re := regexp.MustCompilePOSIX(regexp.QuoteMeta(regexStr))
	result = !re.MatchString(value)
	log.Print(regexStr)
	log.Printf("ContinuousChar: result-%s", strconv.FormatBool(result))
	return
}
