package validator_test

import (
	"gowebapp/validator"
	"strconv"
	"testing"
)

func TestRequired(t *testing.T) {
	/*
		正常系(1文字)
	*/
	// input
	value1 := "a"
	// expected
	expectedValue1 := true
	expectedMsg1 := ""
	// exercise
	actualValue1, actualMsg1 := validator.Required(value1)
	// verify
	if actualValue1 != expectedValue1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue1), strconv.FormatBool(expectedValue1))
	}
	if actualMsg1 != expectedMsg1 {
		t.Errorf("\ngot %s\nexpected %s", actualMsg1, expectedMsg1)
	}

	/*
		正常系(10文字)
	*/
	// input
	value2 := "abcde12345"
	// expected
	expectedValue2 := true
	expectedMsg2 := ""
	// exercise
	actualValue2, actualMsg2 := validator.Required(value2)
	// verify
	if actualValue2 != expectedValue2 {
		t.Errorf("\ngot %s\nexpeted %s\n", strconv.FormatBool(actualValue2), strconv.FormatBool(expectedValue2))
	}
	if actualMsg2 != expectedMsg2 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2, expectedMsg2)
	}

	/*
		未入力
	*/
	// input
	value3 := ""
	// expected
	expectedValue3 := false
	expectedMsg3 := "パスワードが入力されていません"
	// exercise
	actualValue3, actualMsg3 := validator.Required(value3)
	// verify
	if actualValue3 != expectedValue3 {
		t.Errorf("\ngot %s\nexpeted %s\n", strconv.FormatBool(actualValue3), strconv.FormatBool(expectedValue3))
	}
	if actualMsg3 != expectedMsg3 {
		t.Errorf("\ngot %s\nexpeted %s\n", actualMsg3, expectedMsg3)
	}
}

func TestCharType(t *testing.T) {
	/*
		正常系(半角数字のみ)
	*/
	// input
	value1 := "1234567890"
	// expected
	expectedValue1 := true
	expectedMsg1 := ""
	// exercise
	actualValue1, actualMsg1 := validator.CharType(value1)
	// verify
	if actualValue1 != expectedValue1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue1), strconv.FormatBool(expectedValue1))
	}
	if actualMsg1 != expectedMsg1 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg1, expectedMsg1)
	}
	/*
		正常系(半角英字のみ)
	*/
	// input
	value2 := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// expected
	expectedValue2 := true
	expectedMsg2 := ""
	// exercise
	actualValue2, actualMsg2 := validator.CharType(value2)
	// verify
	if actualValue2 != expectedValue2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue2), strconv.FormatBool(expectedValue2))
	}
	if actualMsg2 != expectedMsg2 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2, expectedMsg2)
	}

	/*
		正常系(半角記号のみ)
	*/
	// input
	value3 := "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	// expected
	expectedValue3 := true
	expectedMsg3 := ""
	// exercise
	actualValue3, actualMsg3 := validator.CharType(value3)
	// verify
	if actualValue3 != expectedValue3 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue3), strconv.FormatBool(expectedValue3))
	}
	if actualMsg3 != expectedMsg3 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg3, expectedMsg3)
	}

	/**
	正常系(半角英字、数字、記号混合)
	*/
	// input
	value4 := value1 + value2 + value3
	// expected
	expectedValue4 := true
	expectedMsg4 := ""
	// exercise
	actualValue4, actualMsg4 := validator.CharType(value4)
	// verify
	if actualValue4 != expectedValue4 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue4), strconv.FormatBool(expectedValue4))
	}
	if actualMsg4 != expectedMsg4 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg4, expectedMsg4)
	}

	/*
		全角数字
	*/
	// input
	value5 := "０１２３４５６７８９"
	// expected
	expectedValue5 := false
	expectedMsg5 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue5, actualMsg5 := validator.CharType(value5)
	// verify
	if actualValue5 != expectedValue5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue5), strconv.FormatBool(expectedValue5))
	}
	if actualMsg5 != expectedMsg5 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg5, expectedMsg5)
	}

	/*
		全角英字
	*/
	// input
	value6 := "ａｂｃｄｅｆｇｈｉｊｋｌｍｎｏｐｑｒｓｔｕｖｗｘｙｚＡＢＣＤＥＦＧＨＩＪＫＬＭＮＯＰＱＲＳＴＵＶＷＸＹＺ"
	// expected
	expectedValue6 := false
	expectedMsg6 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue6, actualMsg6 := validator.CharType(value6)
	// verify
	if actualValue6 != expectedValue6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue6), strconv.FormatBool(expectedValue6))
	}
	if actualMsg6 != expectedMsg6 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg6, expectedMsg6)
	}

	/*
		全角記号
	*/
	// input
	value7 := "！”＃＄％＆’（）＊＋,－．／：；＜＝＞？＠［￥］＾＿’｛｜｝～"
	// expected
	expectedValue7 := false
	expectedMsg7 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue7, actualMsg7 := validator.CharType(value7)
	// verify
	if actualValue7 != expectedValue7 {
		t.Errorf("\ngot %s\nexepcted %s\n", strconv.FormatBool(actualValue7), strconv.FormatBool(expectedValue7))
	}
	if actualMsg7 != expectedMsg7 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg7, expectedMsg7)
	}

	/*
		半角カナ
	*/
	// input
	value8 := "ｱｲｳｴｵｶｷｸｹｺﾜｦﾝ"
	// expected
	expectedValue8 := false
	expectedMsg8 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue8, actualMsg8 := validator.CharType(value8)
	// verify
	if actualValue8 != expectedValue8 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue8), strconv.FormatBool(expectedValue8))
	}
	if actualMsg8 != expectedMsg8 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg8, expectedMsg8)
	}

	/*
		全角かな
	*/
	// input
	value9 := "あいうえおかきくけこわをん"
	// expected
	expectedValue9 := false
	expectedMsg9 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue9, actualMsg9 := validator.CharType(value9)
	// verify
	if actualValue9 != expectedValue9 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue9), strconv.FormatBool(expectedValue9))
	}
	if actualMsg9 != expectedMsg9 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg9, actualMsg9)
	}

	/*
		漢字
	*/
	// input
	value10 := "漢字"
	// expected
	expectedValue10 := false
	expectedMsg10 := "パスワードに使用できる文字は半角英字、半角数字、半角記号のみです"
	// exercise
	actualValue10, actualMsg10 := validator.CharType(value10)
	// verify
	if actualValue10 != expectedValue10 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue10), strconv.FormatBool(expectedValue10))
	}
	if actualMsg10 != expectedMsg10 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg10, expectedMsg10)
	}

	/*
		未入力
	*/
	// input
	value11 := ""
	// expected
	expectedValue11 := true
	expectedMsg11 := ""
	// exercise
	actualValue11, actualMsg11 := validator.CharType(value11)
	// verify
	if actualValue11 != expectedValue11 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue11), strconv.FormatBool(expectedValue11))
	}
	if actualMsg11 != expectedMsg11 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg11, expectedMsg11)
	}
}

func TestMaxLength(t *testing.T) {
	/*
		正常系(0文字)
	*/
	// input
	value1 := ""
	// expected
	expectedValue1 := true
	expectedMsg1 := ""
	// exercise
	actualValue1, actualMsg1 := validator.MaxLength(value1)
	// verify
	if actualValue1 != expectedValue1 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue1), strconv.FormatBool(expectedValue1))
	}
	if actualMsg1 != expectedMsg1 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg1, expectedMsg1)
	}

	/*
		正常系(1文字)
	*/
	// input
	value2 := "a"
	// expected
	expectedValue2 := true
	expectedMsg2 := ""
	// exercise
	actualValue2, actualMsg2 := validator.MaxLength(value2)
	// verify
	if actualValue2 != expectedValue2 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue2), strconv.FormatBool(expectedValue2))
	}
	if actualMsg2 != expectedMsg2 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg2, expectedMsg2)
	}

	/*
		正常系(255文字)
	*/
	// input
	value3 := "abcdefghij1234567890abcdefghij1234567890abcdefghij1234567890abcdefghij1234567890abcdefghij1234567890"
	value3 += value3
	value3 += "abcdefghij1234567890abcdefghij1234567890abcdefghij12345"
	if len(value3) != 255 {
		t.Errorf("前提条件の誤り")
	}
	// expeted
	expectedValue3 := true
	expectedMsg3 := ""
	// exercise
	actualValue3, actualMsg3 := validator.MaxLength(value3)
	// verify
	if actualValue3 != expectedValue3 {
		t.Errorf("\ngot %s\nexpeted %s\n", strconv.FormatBool(actualValue3), strconv.FormatBool(expectedValue3))
	}
	if actualMsg3 != expectedMsg3 {
		t.Errorf("\ngot %s\nexpeted %s\n", actualMsg3, expectedMsg3)
	}

	/*
		正常系(256文字)
	*/
	// input
	value3 += "a"
	if len(value3) != 256 {
		t.Errorf("前提条件の誤り")
	}
	// expeted
	expectedValue4 := true
	expectedMsg4 := ""
	// exercise
	actualValue4, actualMsg4 := validator.MaxLength(value3)
	// verify
	if actualValue4 != expectedValue4 {
		t.Errorf("\ngot %s\nexpeted %s\n", strconv.FormatBool(actualValue4), strconv.FormatBool(expectedValue4))
	}
	if actualMsg4 != expectedMsg4 {
		t.Errorf("\ngot %s\nexpeted %s\n", actualMsg4, expectedMsg4)
	}

	/*
		超過(257文字)
	*/
	// input
	value3 += "a"
	// expected
	expectedValue5 := false
	expectedMsg5 := "パスワードは256文字以内で入力してください"
	// exercise
	actualValue5, actualMsg5 := validator.MaxLength(value3)
	// verify
	if actualValue5 != expectedValue5 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue5), strconv.FormatBool(expectedValue5))
	}
	if actualMsg5 != expectedMsg5 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg5, expectedMsg5)
	}

	/*
		未入力
	*/
	// input
	value6 := ""
	// expected
	expectedValue6 := true
	expectedMsg6 := ""
	// exercise
	actualValue6, actualMsg6 := validator.MaxLength(value6)
	// verify
	if actualValue6 != expectedValue6 {
		t.Errorf("\ngot %s\nexpected %s\n", strconv.FormatBool(actualValue6), strconv.FormatBool(expectedValue6))
	}
	if actualMsg6 != expectedMsg6 {
		t.Errorf("\ngot %s\nexpected %s\n", actualMsg6, expectedMsg6)
	}
}
