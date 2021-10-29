package main

import (
	"fmt"
	"nToken/nToken"
)

func main(){
	var searchkey string
	fmt.Scanf("%s", &searchkey)
	token := nToken.LexicalAnalyzer(searchkey)
	for _,i := range (token){
		println("id: ", i.Id, "symbol: ", i.Symbol)
	}
	nToken.SyntaxDirectedTranslator(token)
}
