package checker

import (
	"gowebapp/jsonReader"
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
func LengthCheck(value string, length int) (ok bool) {
	ok = len(value) >= length
	log.Printf("call LengthCheck: ok{%s}", strconv.FormatBool(ok))
	return
}

/*
ComboUpperLowerCase
対象文字列が大文字英字と小文字英字を含んでいるかを調べます
TODO:肯定先読みができるライブラリを使う（^(?=.*[a-z]+)(?=.*[A-Z]+)[a-zA-Z]+$）
*/
func ComboUpperLowerCase(value string) (ok bool) {
	isUpper, _ := regexp.MatchString(".*[A-Z]+", value)
	isLower, _ := regexp.MatchString(".*[a-z]+", value)
	ok = isUpper && isLower
	log.Printf("call ComboUpperLowerCase: ok{%s}", strconv.FormatBool(ok))
	return
}

/*
ComboCharaType
対象文字列が半角英字、半角数字、半角記号の組み合わせを含んでいるかを調べます
引数で指定した数以上の文字種を含んでいることが、判定結果OKとなるための条件です。
*/
func ComboCharaType(value string, condition int) (ok bool) {
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
	ok = satisfy >= condition
	log.Printf("call ComboCharaType: ok{%s} satisfy{%s}", strconv.FormatBool(ok), strconv.Itoa(satisfy))
	return
}

/*
ContinuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
TODO:正規表現で後方参照を使えるように
*/
/*
func ContinuousChar(value string, condition int) (ok bool) {
	re := regexp.MustCompilePOSIX(regexp.QuoteMeta(".*(.)\\1{" + strconv.Itoa(condition-1) + ",}"))
	match := re.FindString(value)
	ok = !re.MatchString(value)
	log.Print(match)
	log.Printf("call ContinuousChar: ok-%s", strconv.FormatBool(ok))
	return
}
*/

/**
ContinuousChar
対象文字列に同じ文字が連続して含まれていないことを調べます
引数で指定した回数未満の繰り返しに抑えられていることが、判定結果OKとなるための条件です
*/
func ContinuousChar(value string, condition int) (ok bool) {
	ok = true
	for _, c := range value {
		var regex string
		for i := 0; i < condition; i++ {
			regex += string(rune(c))
		}
		ok = -1 == strings.Index(value, regex)
		if !ok {
			break
		}
	}

	log.Printf("call ContinuousChar: ok{%s}", strconv.FormatBool(ok))
	return
}

/*
CommonWords
対象の文字列が一般的な単語を使用していないか調べます
一般的な単語は設定用のファイルから読み込みます
*/
func CommonWords(value string) (ok bool) {

	rows := service.ReadAll()
	for _, row := range rows {
		for _, item := range row {
			hit, _ := regexp.MatchString("(?i)"+item, value)
			ok = !hit
			if !ok {
				log.Printf("call CommonWords: ok{%s}", strconv.FormatBool(ok))
				return
			}
		}
	}
	log.Printf("call CommonWords: ok{%s}", strconv.FormatBool(ok))
	return
}

/*
GetSatisfiedCondition
対象の文字列が満足させた条件の数とパスワード強度を上げるための提案メッセージを返します
TODO:理想は引数に与えられた関数だけを呼び出す
func GetSatisfiedCondition(function ...func) (satisfy int)
*/
func GetSatisfiedCondition(password string) (satisfy int, suggestions []string) {
	// 設定ファイルの読み込み
	conf := jsonReader.ReadAll()
	if LengthCheck(password, conf.Length) {
		satisfy++
	} else {
		suggestions = append(suggestions, strconv.Itoa(conf.Length)+"文字以上にしてください")
	}
	if ComboUpperLowerCase(password) {
		satisfy++
	} else {
		suggestions = append(suggestions, "大文字と小文字の組み合わせにしてください")
	}
	if ComboCharaType(password, conf.NumComboChar) {
		satisfy++
	} else {
		suggestions = append(suggestions, "文字種（英字、数字、記号等）を組み合わせましょう")
	}
	if ContinuousChar(password, conf.NumContinuousChar) {
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
