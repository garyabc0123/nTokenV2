package nToken

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/emirpasic/gods/stacks/arraystack"
	"strconv"
)

func Compiler(searchkey string) ([]*ParseTree, error) {
	token := LexicalAnalyzer(searchkey)
	var computeTupleTree []*ParseTree
	err, computeTupleTree := SyntaxDirectedTranslator(token)
	return computeTupleTree, err
}

/**
 * convert string to token stream
 * @param str input string
 * @return ret token stream
 */
func LexicalAnalyzer(str string) (ret []TokenStream) {
	var strRune []rune = []rune(str)
	var buffer bytes.Buffer
	var id int = 0
	for it := 0; it < len(strRune); it++ {
		switch strRune[it] {
		case '\\':
			buffer.WriteRune(strRune[it+1])
			it++

		case '$', '%', '|', '!', '^', '[', ']', '{', '}':
			if buffer.Len() != 0 {
				temp := buffer.String()
				ret = append(ret, TokenStream{id, temp})
				id++
				buffer.Reset()
			}
			ret = append(ret, TokenStream{id, string(strRune[it])})
			id++

		case ' ', '\n', '\t':
			if buffer.Len() != 0 {
				temp := buffer.String()
				ret = append(ret, TokenStream{id, temp})
				id++
				buffer.Reset()
			}

		default:
			buffer.WriteRune(strRune[it])

		}
	}
	if buffer.Len() != 0 {
		temp := buffer.String()
		ret = append(ret, TokenStream{id, temp})
	}
	return
}

/**
 * convert token stream from in-order to pre-order
 * than create treee
 * @param tokenStream
 * @return error
 * @return a list saving pointer to parse tree root
 */
func SyntaxDirectedTranslator(tokenStream []TokenStream) (error, []*ParseTree) {

	var nowState int = 0
	var curlyBegin, curlyEnd int
	var squareBegin, squareEnd int
	//0 - not thing, check { is exist
	//1 - {, check } is exist
	//2 - }, check [ is exist
	//3 - [, check ] is exist
	//4 - ], reset
	stack := arraystack.New()
	var computeTupleTree []*ParseTree

	for it := 0; it < len(tokenStream); it++ {
		switch nowState {
		case 0:
			if len(tokenStream[it].Symbol) != 1 {
				continue
			}
			if tokenStream[it].Symbol[0] == curlyBracketLeft {
				curlyBegin = it
				nowState++
				stack.Push(curlyBracketLeft)
			}
		case 1:
			if len(tokenStream[it].Symbol) != 1 {
				continue
			}
			if tokenStream[it].Symbol[0] == curlyBracketLeft {
				stack.Push(curlyBracketLeft)
				continue
			}
			if tokenStream[it].Symbol[0] == curlyBracketRight {
				if !stack.Empty() {
					stack.Pop()
					if stack.Empty() {
						curlyEnd = it
						nowState++
					}
				} else {
					return errors.New("Synatex Error, loss {"), nil
				}
			}
		case 2:
			if len(tokenStream[it].Symbol) != 1 {
				continue
			}
			if tokenStream[it].Symbol[0] == squareBracketLeft {
				squareBegin = it
				nowState++
			}
		case 3:
			if len(tokenStream[it].Symbol) != 1 {
				continue
			}
			if tokenStream[it].Symbol[0] == squareBracketRight {
				squareEnd = it
				nowState = 0

				computeTupleTree = append(computeTupleTree, tokenStream2Tree(tokenStream[curlyBegin+1:curlyEnd]))

				distStr := ""

				for it2 := squareBegin + 1; it2 < squareEnd; it2++ {
					distStr += tokenStream[it2].Symbol
				}
				dist, err := strconv.Atoi(distStr)
				if err != nil {
					return err, nil
				}
				fmt.Println(curlyBegin, curlyEnd, dist)
			}
		}
	}
	return nil, computeTupleTree
}

func infixToPrefix(input []TokenStream) (output []TokenStream) {
	var stack *arraystack.Stack = arraystack.New()
	for it := len(input) - 1; it >= 0; it-- {
		if len(input[it].Symbol) != 1 {
			output = append([]TokenStream{input[it]}, output...)
		} else {
			switch input[it].Symbol[0] {
			case curlyBracketRight:
				stack.Push(input[it])
			case curlyBracketLeft:

				for !stack.Empty() {
					temp, _ := stack.Pop()
					if temp.(TokenStream).Symbol[0] == curlyBracketRight {
						break
					} else {
						output = append([]TokenStream{temp.(TokenStream)}, output...)
					}
				}
			case verticalBar, caret:
				if stack.Empty() {
					stack.Push(input[it])
					continue
				}
				for !stack.Empty() {
					temp, _ := stack.Pop()
					if temp.(TokenStream).Symbol[0] == curlyBracketRight {
						stack.Push(temp)
						break
					} else if operatorPriority(SymbolTable(temp.(TokenStream).Symbol[0])) < operatorPriority(SymbolTable(input[it].Symbol[0])) {
						output = append([]TokenStream{temp.(TokenStream)}, output...)
					} else {
						stack.Push(temp)
						break
					}
				}
				stack.Push(input[it])

			case percentSign, exclamationMark, dollerSign:
				output = append([]TokenStream{input[it]}, output...)

			default:

				output = append([]TokenStream{input[it]}, output...)

			}
		}
	}

	for !stack.Empty() {
		temp, _ := stack.Pop()
		output = append([]TokenStream{temp.(TokenStream)}, output...)
	}

	return
}

func prefixToParseTree(input *[]TokenStream, begin int, size int, me *ParseTree) (next int) {
	for it := begin; it < begin+size && it < len(*input); it++ {
		switch (*input)[it].Symbol[0] {
		case dollerSign, percentSign, exclamationMark:
			me.Token = (*input)[it]
			me.Left = new(ParseTree)
			it = prefixToParseTree(input, it+1, 1, me.Left)
			next = it
		case verticalBar, caret:
			me.Token = (*input)[it]
			me.Left = new(ParseTree)
			me.Right = new(ParseTree)
			it = prefixToParseTree(input, it+1, 1, me.Left)
			it = prefixToParseTree(input, it+1, 1, me.Right)
			next = it

		default:
			me.Token = (*input)[it]
			next = it

		}
	}

	return
}

func tokenStream2Tree(tokenStream []TokenStream) *ParseTree {
	tokenStream = infixToPrefix(tokenStream)
	for _, itt := range tokenStream {
		print(itt.Symbol, " ")
	}
	var treeHead *ParseTree = new(ParseTree)
	treeHead.Token.Id = -1
	prefixToParseTree(&tokenStream, 0, len(tokenStream), treeHead)

	//TODO convert to nfa

	return treeHead
}
