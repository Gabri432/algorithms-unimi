package algorithmsunimi

import "testing"

func TestParenthesisChecker(t *testing.T) {
	parenthesisChecker("{([(<{}[<>[]}>{[]{[(<()>")
	parenthesisChecker("[[<[([]))<([[{}[[()]]]")
	parenthesisChecker("[{[{({}]{}}([{[{{{}}([]")
	parenthesisChecker("[<(<(<(<{}))><([]([]()")
	parenthesisChecker("<{([([[(<>()){}]>(<<{{")
}
