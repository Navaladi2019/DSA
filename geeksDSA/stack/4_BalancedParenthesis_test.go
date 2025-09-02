package stack

import "testing"

func Test_IsBalancedParenthesis(t *testing.T) {

	got := IsBalancedParanthesis_SImple("([])")

	if got != true {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParanthesis_SImple("((())")

	if got != false {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParanthesis_SImple("()")

	if got != true {
		t.Error("Has error in parenthesis")
	}

	got = IsBalancedParanthesis_SImple("{}{}([()])")

	if got != true {
		t.Error("Has error in parenthesis")
	}
}
