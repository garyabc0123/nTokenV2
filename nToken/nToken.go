package nToken

type TokenStream struct {

	Id int
	Symbol string
}

type SymbolTable byte

const (
	dollerSign 				= '$'
	percentSign				= '%'
	verticalBar 			= '|'
	exclamationMark			= '!'
	caret					= '^'
	squareBracketLeft		= '['
	squareBracketRight		= ']'
	curlyBracketLeft		= '{'
	curlyBracketRight		= '}'
)
func operatorPriority(in SymbolTable)int{
	switch in {
	case squareBracketLeft, squareBracketRight, curlyBracketLeft, curlyBracketRight:
		return 3
	case dollerSign, percentSign:
		return 4
	case exclamationMark:
		return 5
	case caret:
		return 6
	case verticalBar:
		return 7
	default:
		return 100
	}
}

type parseTree struct{
	token TokenStream
	left *parseTree
	right *parseTree

}