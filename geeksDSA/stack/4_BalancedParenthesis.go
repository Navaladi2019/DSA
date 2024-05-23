package stack

func IsBalancedParenthesis(str string) bool {

	r := []rune(str)

	if len(r)%2 != 0 {
		return false
	}

	sta := LLStack[rune]{}

	for i := 0; i < len(r); i++ {

		lastval, ok := sta.Peek()

		isclosing := ismatch(lastval, r[i])

		if ok && isclosing {
			sta.Pop()
		} else {
			sta.Push(r[i])
		}

	}

	return sta.isEmpty()
}

func ismatch(open rune, close rune) bool {

	switch open {

	case '{':

		if close == '}' {
			return true
		}
		return false

	case '(':

		if close == ')' {
			return true
		}
		return false
	case '[':

		if close == ']' {
			return true
		}
		return false

	default:
		return false

	}
}
