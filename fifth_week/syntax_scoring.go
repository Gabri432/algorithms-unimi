package algorithmsunimi

//https://adventofcode.com/2021/day/10

import (
	"fmt"
	"strings"
)

func parenthesisChecker(textLine string) {
	textLine = removeClosingParenthesis(textLine)
	for i := 1; i-1 < len(textLine); i++ {
		currentParenthesis := string(textLine[i])
		nextParenthesis := string(textLine[i+1])
		if isOpeningParenthesis(currentParenthesis) && isClosingParenthesis(nextParenthesis) && !areClosing(currentParenthesis, nextParenthesis) { //case corrupted
			printError(currentParenthesis, nextParenthesis)
			return
		}
	}
}

// Removes every "()", "[]", "<>", "{}" (opening and closing parenthesis) from the string.
func removeClosingParenthesis(textLine string) string {
	textLine = strings.ReplaceAll(textLine, "()", "")
	textLine = strings.ReplaceAll(textLine, "[]", "")
	textLine = strings.ReplaceAll(textLine, "{}", "")
	newTextLine := strings.ReplaceAll(textLine, "<>", "")
	if !strings.Contains(newTextLine, "()") && !strings.Contains(newTextLine, "[]") && !strings.Contains(newTextLine, "{}") && !strings.Contains(textLine, "<>") {
		return newTextLine
	}
	return removeClosingParenthesis(newTextLine)
}

func isOpeningParenthesis(p string) bool {
	switch p {
	case "{", "<", "[", "(":
		return true
	default:
		return false
	}
}

func isClosingParenthesis(p string) bool {
	switch p {
	case "}", ">", "]", ")":
		return true
	default:
		return false
	}
}

func areClosing(p1, p2 string) bool {
	switch {
	case p1 == "[" && p2 == "]", p1 == "{" && p2 == "}", p1 == "<" && p2 == ">", p1 == "(" && p2 == ")": // cases [], <>, (), []
		return true
	default:
		return false
	}
}

func printError(p1, p2 string) {
	switch p1 {
	case "[":
		fmt.Printf("Expected ], but found %s instead.\n", p2)
	case "{":
		fmt.Printf("Expected }, but found %s instead.\n", p2)
	case "(":
		fmt.Printf("Expected ), but found %s instead.\n", p2)
	case "<":
		fmt.Printf("Expected >, but found %s instead.\n", p2)
	}
}
