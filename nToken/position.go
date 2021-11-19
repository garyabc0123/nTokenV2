package nToken

import (
	"errors"
	"github.com/viney-shih/goroutines"
	"runtime"
	"strconv"
)

func GetConformPosition(document Document, tree []*ParseTree) []DocumentPositionList {
	var position []DocumentPositionList = make([]DocumentPositionList, len(tree)) //search keyword pair - sentence id - serial id
	for i, _ := range tree {
		position[i] = make(DocumentPositionList, len(document))
	}
	batch := goroutines.NewBatch(runtime.NumCPU()*2, goroutines.WithBatchSize(len(tree)))
	for id, _ := range tree {
		idx := id
		batch.Queue(func() (interface{}, error) {
			myTree := tree[idx]
			for ix, _ := range document {
				for iy, _ := range *(document[ix]) {
					if b, e := Treetraversal(myTree, *(*((document)[ix]))[iy]); b && e == nil {
						position[idx][ix] = append(position[idx][ix], Position(iy))
					} else if e != nil {
						return nil, e
					}
				}
			}

			return nil, nil
		})
	}

	batch.QueueComplete()
	batch.Close()
	/*for pairId := range position {
		println("ID:", pairId)

		for senId := range position[pairId] {
			println("\tsentence: ", senId)
			print("\t\t\t")
			for i := range position[pairId][senId] {
				print(position[pairId][senId][i], " ")
			}
			println()
		}

	}*/
	return position
}

func Treetraversal(node *ParseTree, key WordAndPartOfSpeechPair) (bool, error) {
	switch {
	////////////////////////////////////////////////////////////////////////////
	case len(node.Token.Symbol) == 1 && node.Token.Symbol[0] == dollerSign:
		if node.Left == nil {
			return false, errors.New("error tree at " + strconv.Itoa(node.Token.Id))
		}
		return node.Left.Token.Symbol == string(key.Word), nil
	////////////////////////////////////////////////////////////////////////////
	case len(node.Token.Symbol) == 1 && node.Token.Symbol[0] == percentSign:
		if node.Left == nil {
			return false, errors.New("error tree at " + strconv.Itoa(node.Token.Id))
		}
		temp, err := strconv.Atoi(node.Left.Token.Symbol)
		if err != nil {
			return false, errors.New("error part of speech number" + node.Left.Token.Symbol)
		}
		return temp == key.PartOfSpeech, nil

	////////////////////////////////////////////////////////////////////////////
	case len(node.Token.Symbol) == 1 && node.Token.Symbol[0] == verticalBar:
		if node.Left == nil || node.Right == nil {
			return false, errors.New("error tree at " + strconv.Itoa(node.Token.Id))
		}
		var left, right bool
		left, err := Treetraversal(node.Left, key)
		if err != nil {
			return false, err
		}
		right, err = Treetraversal(node.Right, key)
		if err != nil {
			return false, err
		}
		return left || right, nil

	////////////////////////////////////////////////////////////////////////////
	case len(node.Token.Symbol) == 1 && node.Token.Symbol[0] == exclamationMark:
		if node.Left == nil {
			return false, errors.New("error tree at " + strconv.Itoa(node.Token.Id))
		}
		left, err := Treetraversal(node.Left, key)
		return !left, err

	////////////////////////////////////////////////////////////////////////////
	case len(node.Token.Symbol) == 1 && node.Token.Symbol[0] == caret:
		if node.Left == nil || node.Right == nil {
			return false, errors.New("error tree at " + strconv.Itoa(node.Token.Id))
		}
		var left, right bool
		left, err := Treetraversal(node.Left, key)
		if err != nil {
			return false, err
		}
		right, err = Treetraversal(node.Right, key)
		if err != nil {
			return false, err
		}
		return left && right, nil

	}
	return false, nil
}
