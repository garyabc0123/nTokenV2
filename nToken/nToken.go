package nToken

/*
 * Search Queue Token Stream
 */
type TokenStream struct {
	Id     int
	Symbol string
}

/**
 * convert Symbol to symbol priority
 * number lower have high priority
 * @param in input Symbol
 * @return symbol priority
 */
func operatorPriority(in SymbolTable) int {
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

type ParseTree struct {
	Token TokenStream
	Left  *ParseTree
	Right *ParseTree
}

/**
 * Symbol List
 */
type SymbolTable byte

const (
	dollerSign         = '$'
	percentSign        = '%'
	verticalBar        = '|'
	exclamationMark    = '!'
	caret              = '^'
	squareBracketLeft  = '['
	squareBracketRight = ']'
	curlyBracketLeft   = '{'
	curlyBracketRight  = '}'
)

type WordAndPartOfSpeechPair struct {
	Id          int
	Word        []rune
	PartOfSeech int
}

type Sentence []*WordAndPartOfSpeechPair
type Document []*Sentence
