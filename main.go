package main

import (
	"nToken/nToken"
)

func main() {
	//defer profile.Start().Stop()

	temp, err := nToken.ReadData("solr.dat")
	if err != nil {
		panic(err)
	}

	searchKey, err := nToken.ReadSearchQueue("search.txt")
	if err != nil {
		panic(err)
	}
	computeTupleTree, distList, err := nToken.Compiler(searchKey)
	if err != nil {
		panic(err)
	}
	println(distList[0])

	nToken.GetConformPosition(temp, computeTupleTree)

}
