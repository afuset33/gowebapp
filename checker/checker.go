package checker

import (
	"gowebapp/service"
	"log"
	"regexp"
	"strconv"
	"strings"
)

/*
LengthCheck
対象文字列が指定桁数以上であるかを調べます
引数で指定した以上の桁数であることが、判定結果OKとなるための条件です
*/
func LengthCheck(value string, length int) (result bool) {
	result = len(value) >= length
	log.Printf("call LengthCheck: result{%s}", strconv.FormatBool(result))
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
	log.Printf("call ComboUpperLowerCase: result{%s}", strconv.FormatBool(result))
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
	log.Printf("call ComboCharaType: result{%s} satisfy{%s}", strconv.FormatBool(result), strconv.Itoa(satisfy))
	return
}

/*
ContinuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
TODO:正規表現で後方参照を使えるように
*/
/*
func ContinuousChar(value string, condition int) (result bool) {
	re := regexp.MustCompilePOSIX(regexp.QuoteMeta(".*(.)\\1{" + strconv.Itoa(condition-1) + ",}"))
	match := re.FindString(value)
	result = !re.MatchString(value)
	log.Print(match)
	log.Printf("call ContinuousChar: result-%s", strconv.FormatBool(result))
	return
}
*/

/**
ContinuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
*/
func ContinuousChar(value string, condition int) (result bool) {

	for _, c := range value {
		var regex string
		for i := 0; i < condition; i++ {
			regex += string(rune(c))
		}
		result = -1 == strings.Index(value, regex)
		if !result {
			log.Printf("call ContinuousChar: result{%s}", strconv.FormatBool(result))
			return
		}
	}

	log.Printf("call ContinuousChar: result{%s}", strconv.FormatBool(result))
	return
}

/*
CommonWords
対象の文字列が一般的な単語を使用していないか調べます
一般的な単語は設定用のファイルから読み込みます
*/
func CommonWords(value string) (result bool) {

	rows := service.ReadAll()
	for _, row := range rows {
		for _, item := range row {
			result = -1 == strings.Index(value, item)
			if !result {
				log.Printf("call CommonWords: result{%s}", strconv.FormatBool(result))
				return
			}
		}
	}
	log.Printf("call CommonWords: result{%s}", strconv.FormatBool(result))
	return
}

/*
GetSatisfiedCondition
対象の文字列が満足させた条件の数とパスワード強度を上げるための提案メッセージを返します
TODO:理想は引数に与えられた関数だけを呼び出す
func GetSatisfiedCondition(function ...func) (satisfy int)
*/
func GetSatisfiedCondition(password string) (satisfy int, suggestions []string) {
	if LengthCheck(password, 8) {
		satisfy++
	} else {
		suggestions = append(suggestions, "8文字以上にしてください")
	}
	if ComboUpperLowerCase(password) {
		satisfy++
	} else {
		suggestions = append(suggestions, "大文字と小文字の組み合わせにしてください")
	}
	if ComboCharaType(password, 2) {
		satisfy++
	} else {
		suggestions = append(suggestions, "文字種（英字、数字、記号等）を組み合わせましょう")
	}
	if ContinuousChar(password, 3) {
		satisfy++
	} else {
		suggestions = append(suggestions, "同じ文字を連続して使用するのはやめましょう")
	}
	if CommonWords(password) {
		satisfy++
	} else {
		suggestions = append(suggestions, "一般的な単語を使用するのはやめましょう")
	}
	log.Print("call GetSatisfiedNum: satisfy{" + strconv.Itoa(satisfy) + "}")
	return
}
