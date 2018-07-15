package checker

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
lengthCheck
対象文字列が指定桁数以上であるかを調べます
引数で指定した以上の桁数であることが、判定結果OKとなるための条件です
*/
func lengthCheck(value string, length int) (result bool) {
	result = len(value) >= length
	log.Printf("call lengthCheck: result-%s", strconv.FormatBool(result))
	return
}

/*
comboUpperLowerCase
対象文字列が大文字英字と小文字英字を含んでいるかを調べます
TODO:肯定先読みができるライブラリを使う（^(?=.*[a-z]+)(?=.*[A-Z]+)[a-zA-Z]+$）
*/
func comboUpperLowerCase(value string) (result bool) {
	isUpper, _ := regexp.MatchString(".*[A-Z]+", value)
	isLower, _ := regexp.MatchString(".*[a-z]+", value)
	result = isUpper && isLower
	log.Printf("call comboUpperLowerCase: result-%s", strconv.FormatBool(result))
	return
}

/*
comboCharaType
対象文字列が半角英字、半角数字、半角記号の組み合わせを含んでいるかを調べます
引数で指定した数以上の文字種を含んでいることが、判定結果OKとなるための条件です。
*/
func comboCharaType(value string, condition int) (result bool) {
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
	log.Printf("call comboCharaType: result-%s satisfy-%s", strconv.FormatBool(result), strconv.Itoa(satisfy))
	return
}

/*
continuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
TODO:正規表現で後方参照を使えるように
*/
/*
func continuousChar(value string, condition int) (result bool) {
	re := regexp.MustCompilePOSIX(regexp.QuoteMeta(".*(.)\\1{" + strconv.Itoa(condition-1) + ",}"))
	match := re.FindString(value)
	result = !re.MatchString(value)
	log.Print(match)
	log.Printf("call continuousChar: result-%s", strconv.FormatBool(result))
	return
}
*/

/**
continuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
*/
func continuousChar(value string, condition int) (result bool) {

	for _, c := range value {
		var regex string
		for i := 0; i < condition; i++ {
			regex += string(rune(c))
		}
		result = -1 == strings.Index(value, regex)
		if !result {
			break
		}
	}

	log.Printf("call continuousChar: result-%s", strconv.FormatBool(result))
	return
}

/*
GetSatisfiedNums
対象の文字列が満足させた条件の数を返します
TODO:理想は引数に与えられた関数だけを呼び出す
func GetSatisfiedNum(function ...func) (satisfy int)
*/
func GetSatisfiedNum(password string) (satisfy int) {
	if lengthCheck(password, 8) {
		satisfy++
	}
	if comboUpperLowerCase(password) {
		satisfy++
	}
	if comboCharaType(password, 2) {
		satisfy++
	}
	if continuousChar(password, 3) {
		satisfy++
	}
	log.Print("call GetSatisfiedNum: satisfy-" + strconv.Itoa(satisfy))
	return
}
