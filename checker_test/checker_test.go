package checker_test

import (
	"gowebapp/checker"
	"strconv"
	"testing"
)

func TestLengthCheck(t *testing.T) {
	/*
		input 7文字
		condition 8文字
	*/
	// input 7文字
	value1 := "abcdefg"
	condition1 := 8
	// expected
	expected1 := false
	// verify
	actual1 := checker.LengthCheck(value1, condition1)
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		input 8文字
		condition 8文字
	*/
	// input
	value2 := "abcdefgh"
	condition2 := 8
	// expected
	expected2 := true
	// exercise
	actual2 := checker.LengthCheck(value2, condition2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/**
	input 9文字
	condition 8文字
	*/
	// input
	value3 := "abcdefghi"
	condition3 := 8
	// expected
	expected3 := true
	// exercise
	actual3 := checker.LengthCheck(value3, condition3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/**
	input 9文字
	condition 10文字
	*/
	// input
	value4 := "abcdefghi"
	condition4 := 10
	// expected
	expected4 := false
	// exercise
	actual4 := checker.LengthCheck(value4, condition4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		input 0文字
		condition 8文字
	*/
	// input
	value5 := ""
	condition5 := 8
	// expected
	expected5 := false
	// exercise
	actual5 := checker.LengthCheck(value5, condition5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		input 1文字
		condition 0文字
	*/
	// input
	value6 := "a"
	condition6 := 0
	// expected
	expected6 := true
	// exercise
	actual6 := checker.LengthCheck(value6, condition6)
	// verfiy
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		input 0文字
		condition 0文字
	*/
	// input
	value7 := ""
	condition7 := 0
	// expected
	expected7 := true
	// exercise
	actual7 := checker.LengthCheck(value7, condition7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}
}

func TestComboUpperLowerCase(t *testing.T) {
	/*
		小文字のみ
	*/
	// input
	value1 := "abcdefg"
	// expected
	expected1 := false
	// exercise
	actual1 := checker.ComboUpperLowerCase(value1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		大文字のみ
	*/
	// input
	value2 := "ABCDEFG"
	// expected
	expected2 := false
	// exercise
	actual2 := checker.ComboUpperLowerCase(value2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		大文字小文字混合
	*/
	// input
	value3 := "abcdABCD"
	// expected
	expected3 := true
	// exercise
	actual3 := checker.ComboUpperLowerCase(value3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		全角英字小文字大文字混合
	*/
	// input
	value4 := "ａｂｃｄＡＢＣＤ"
	// expected
	expected4 := false
	// exercise
	actual4 := checker.ComboUpperLowerCase(value4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		半角数字
	*/
	// input
	value5 := "1234567890"
	// expected
	expected5 := false
	// exercise
	actual5 := checker.ComboUpperLowerCase(value5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		半角記号
	*/
	// input
	value6 := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	// expected
	expected6 := false
	// exercise
	actual6 := checker.ComboUpperLowerCase(value6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		半角英字小文字大文字、半角数字、半角記号混合
	*/
	// input
	value7 := "abcABC01234!.:@;"
	// expected
	expected7 := true
	// exercise
	actual7 := checker.ComboUpperLowerCase(value7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		半角英字小文字大文字、全角英字混合
	*/
	// input
	value8 := "abcABCａｂｃＡＢＣ"
	// expected
	expected8 := true
	// exercise
	actual8 := checker.ComboUpperLowerCase(value8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		未入力
	*/
	// input
	value9 := ""
	// expected
	expected9 := false
	// exercise
	actual9 := checker.ComboUpperLowerCase(value9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}
}

func TestComboCharType(t *testing.T) {
	/*
		半角英字、半角数字
		condition 2
	*/
	// input
	value1 := "abcde12345"
	condition1 := 2
	// expected
	expected1 := true
	// exercise
	actual1 := checker.ComboCharaType(value1, condition1)
	// verify
	if actual1 != expected1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual1), strconv.FormatBool(expected1))
	}

	/*
		半角英字、半角記号
		condition 2
	*/
	// input
	value2 := "abcde@:[]"
	condition2 := 2
	// expected
	expected2 := true
	// exercise
	actual2 := checker.ComboCharaType(value2, condition2)
	// verify
	if actual2 != expected2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual2), strconv.FormatBool(expected2))
	}

	/*
		半角記号、半角数字
		condition 2
	*/
	// input
	value3 := "[@p:;1234567890"
	condition3 := 2
	// expected
	expected3 := true
	// exercise
	actual3 := checker.ComboCharaType(value3, condition3)
	// verify
	if actual3 != expected3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual3), strconv.FormatBool(expected3))
	}

	/*
		半角英字、半角数字、半角記号
		condition 2
	*/
	// input
	value4 := "abcABC0123456789[@];,'()"
	condition4 := 2
	// expected
	expected4 := true
	// exercise
	actual4 := checker.ComboCharaType(value4, condition4)
	// verify
	if actual4 != expected4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual4), strconv.FormatBool(expected4))
	}

	/*
		半角英字、半角カナ
		condition 2
	*/
	// input
	value5 := "abcdeｱｲｳｴｵﾜｦﾝ"
	condition5 := 2
	// expected
	expected5 := false
	// exercise
	actual5 := checker.ComboCharaType(value5, condition5)
	// verify
	if actual5 != expected5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual5), strconv.FormatBool(expected5))
	}

	/*
		半角英字、半角数字、半角記号、半角カナ
		condition 2
	*/
	// input
	value6 := "abcdeABCDE1234567890:;@'(&%ｱｲｳｴｵﾜｦﾝ"
	condition6 := 2
	// expected
	expected6 := true
	// exercise
	actual6 := checker.ComboCharaType(value6, condition6)
	// verify
	if actual6 != expected6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual6), strconv.FormatBool(expected6))
	}

	/*
		半角英字
		condition 2
	*/
	// input
	value7 := "abcdeABCDE"
	condition7 := 2
	// expected
	expected7 := false
	// exercise
	actual7 := checker.ComboCharaType(value7, condition7)
	// verify
	if actual7 != expected7 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual7), strconv.FormatBool(expected7))
	}

	/*
		半角数字
		condition 2
	*/
	// input
	value8 := "0123456789"
	condition8 := 2
	// expected
	expected8 := false
	// exercise
	actual8 := checker.ComboCharaType(value8, condition8)
	// verify
	if actual8 != expected8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual8), strconv.FormatBool(expected8))
	}

	/*
		半角記号
		condition 2
	*/
	// input
	value9 := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	condition9 := 2
	// expected
	expected9 := false
	// exercise
	actual9 := checker.ComboCharaType(value9, condition9)
	// verify
	if actual9 != expected9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual9), strconv.FormatBool(expected9))
	}

	/*
		半角英字、半角数字
		condition 3
	*/
	// input
	value10 := "abcde12345"
	condition10 := 3
	// expected
	expected10 := false
	// exercise
	actual10 := checker.ComboCharaType(value10, condition10)
	// verify
	if actual10 != expected10 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual10), strconv.FormatBool(expected10))
	}

	/*
		入力無し
		condition 2
	*/
	// input
	value11 := ""
	condition11 := 2
	// expected
	expected11 := false
	// exercise
	actual11 := checker.ComboCharaType(value11, condition11)
	// verify
	if actual11 != expected11 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual11), strconv.FormatBool(expected11))
	}

	/*
		半角英字
		condition 0
	*/
	// input
	value12 := "abcXYZ"
	condition12 := 0
	// expected
	expected12 := true
	// exercise
	actual12 := checker.ComboCharaType(value12, condition12)
	// verify
	if actual12 != expected12 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual12), strconv.FormatBool(expected12))
	}

	/*
		入力無し
		condition 0
	*/
	// input
	value13 := ""
	condition13 := 0
	// expected
	expected13 := true
	// exercise
	actual13 := checker.ComboCharaType(value13, condition13)
	// verify
	if actual13 != expected13 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actual13), strconv.FormatBool(expected13))
	}
}
