package checker_test

import (
	"gowebapp/checker"
	"testing"
)

func TestlengthCheck(t *testing.T) {
	// value 7文字
	value := "abcdefg"
	// expected 一般単語、文字連続
	const expected int = 2
	// verify
	actual := checker.GetSatisfiedNum(value)
	if actual != expected {
		t.Errorf("got %d\nexpected%d\n", actual, expected)
	}
}
