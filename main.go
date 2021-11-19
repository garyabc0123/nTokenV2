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

	position := nToken.GetConformPosition(temp, computeTupleTree)
	for i := 0; i < len(position); i++ {
		print(computeTupleTree[i].Left.Token.Symbol)
		print("\t")
	}
	println()
	var it int = 0
	for {
		var p bool = false
		for i := 0; i < len(position); i++ {
			if len(position[i][0]) > it {
				print(position[i][0][it])
				print("\t")
				p = true
			} else {
				print("\t")
			}

		}
		if !p {
			break
		}
		println()
		it++
		if it > 1000 {
			break
		}
	}

}
