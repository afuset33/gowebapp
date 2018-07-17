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
