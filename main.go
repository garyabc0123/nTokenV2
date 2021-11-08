package main

import (
	"nToken/nToken"
)

func main() {
	//defer profile.Start().Stop()

	{

		temp, err := nToken.ReadData("/mnt/data/go/solr-ckip.dat")
		if err != nil {
			panic(err)
		}
		println((*(temp[0]))[0].Word)
	}
	{

		computeTupleTree, err := nToken.Compiler("{%1}[0]")
		if err != nil {
			panic(err)
		}
		i := computeTupleTree[0].Token.Symbol
		println(i)
	}

}
