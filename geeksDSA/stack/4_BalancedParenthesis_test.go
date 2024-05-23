package stack

import "testing"

func Test_IsBalancedParenthesis(t *testing.T) {

	got := IsBalancedParenthesis("([])")

	if got != true {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParenthesis("((())")

	if got != false {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParenthesis("()")

	if got != true {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParenthesis("{}{}([()])")

	if got != true {
		t.Error("Has error in parenthesis")
	}
}
