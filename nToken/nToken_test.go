package nToken

import "testing"

func TestNtoken(t *testing.T) {
	temp, err := ReadData("solr_nt.dat")
	if err != nil {
		panic(err)
	}

	searchKey, err := ReadSearchQueue("search_nt.txt")
	if err != nil {
		panic(err)
	}
	computeTupleTree, distList, err := Compiler(searchKey)
	if err != nil {
		panic(err)
	}
	position := GetConformPosition(temp, computeTupleTree)
	ans := NoName(position, distList[:len(distList)-1])

	println()
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

	println()
	println(searchKey)

	for sentenceId := 0; sentenceId < len(ans); sentenceId++ {
		println("sentence id: ", sentenceId)
		if len(ans[sentenceId]) == 0 {
			println("\tno match")
			continue
		}
		for matchId := 0; matchId < len(ans[sentenceId]); matchId++ {
			print("\tmatch id ", matchId, ": ")
			for i := 0; i < len(ans[sentenceId][matchId]); i++ {
				print(ans[sentenceId][matchId][i], ", ")

			}
			println()
		}
	}
}
