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

// ismatch checks if the given pair of runes represents matching opening and closing brackets.
// It returns true if 'open' and 'close' form a valid pair of parentheses, curly braces, or square brackets.
// Supported pairs are: '()', '{}', and '[]'.
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

func IsBalancedParanthesis_SImple(str string) bool {

	r := []rune(str)

	s := ArrStack[rune]{}

	for _, v := range r {
		if v == '(' || v == '[' || v == '{' {
			s.Push(v)
			continue
		}

		if ex, ok := s.Peek(); !(ok && ismatch(ex, v)) {
			return false
		} else {
			s.Pop()
		}

	}

	return s.IsEmpty()
}
