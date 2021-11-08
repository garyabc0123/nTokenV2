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
		//[ ] { }
		return 3
	case dollerSign, percentSign:
		//$ %
		return 4
	case exclamationMark:
		//!
		return 5
	case caret:
		//^
		return 6
	case verticalBar:
		//|
		return 7
	default:
		return 100
	}
}

/**
 * parse tree data struct
 * if only has one child, must put into Left leaf node
 */
type ParseTree struct {
	Token TokenStream
	Left  *ParseTree
	Right *ParseTree
}

/**
 * Symbol List

| 符號 |                說明                 | 優先度(1是最高) |
| ---- |:----------------------------------:|:---------------:|
| \    |   將下一個字元標記為一個特殊字元         |        1        |
| $    |       下一個字元開始為查詢詞           |        4        |
| %    |      下一個字元開始為查詢詞性          |        4        |
| \|   |             or                     |       7        |
| !    |                not                 |       5        |
| ^    |                and                 |       6        |
| {}[] | `{}`裡面是查詢運算，`[]`裡面是間距      |      2         |
| {}   |  括號                               |      3         |

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

/**
 * input data struct
 * must be:
 * word(part of speeh id) word(part of speeh id) word(part of speeh id)...
 */
type WordAndPartOfSpeechPair struct {
	Id           int
	Word         []rune
	PartOfSpeech int
}

type Sentence []*WordAndPartOfSpeechPair
type Document []*Sentence

//position data struct
type Position int
type SentencePositionList []Position
type DocumentPositionList []Sentence
