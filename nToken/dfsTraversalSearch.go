package nToken

import (
	"github.com/emirpasic/gods/stacks/arraystack"
)

func NoName(position []DocumentPositionList, dist DistList) {
	//TODO give name

	for sentenceID := 0; sentenceID < len(position[0]); sentenceID++ {
		for tokenID := 0; tokenID < len(position[0][sentenceID]); tokenID++ {
			dfsRecursive(&position, &dist, 0, PositionArray{sentenceID, append([]Position{}, Position(tokenID))})
		}

	}

}

func dfsRecursive(position *[]DocumentPositionList, dist *DistList, level int, nowPosition PositionArray) (success bool, TokenID []Position) {
	if level == len(*dist) {
		return true, nowPosition.TokenID
	}
	nextTokenId, _ := binarySearchfloor((*position)[level+1][nowPosition.SentenceID], nowPosition.TokenID[len(nowPosition.TokenID)-1]+Position((*dist)[level]))
	if nextTokenId == -1 {
		return false, nil

	}
	for ; nextTokenId >= 0; nextTokenId-- {

		if level > 1 && (*position)[level+1][nowPosition.SentenceID][nextTokenId] <= nowPosition.TokenID[level-1] {
			break
		}
		if (*position)[level+1][nowPosition.SentenceID][nextTokenId] <= nowPosition.TokenID[level] {
			continue
		}

		if x, tokenID := dfsRecursive(
			position,
			dist,
			level+1,
			PositionArray{
				SentenceID: nowPosition.SentenceID,
				TokenID:    append(nowPosition.TokenID, (*position)[level+1][nowPosition.SentenceID][nextTokenId]),
			}); x {
			return true, tokenID
		}
	}

	return false, nil

}
func dfsLoop(position *[]DocumentPositionList, dist *DistList, level int, nowPosition PositionArray) (success bool, TokenID []Position) {
	stack := arraystack.New()
	//stack max size: count
	for it := 0; (*position)[level+1][nowPosition.SentenceID][it] < Position((*dist)[level]); it++ {

	}

	return false, nil
}

//now not binary search :-(
//TODO: let bigO smaller
func binarySearchfloor(array SentencePositionList, num Position) (addr int, data Position) {
	addr = -1
	data = -1
	for it, d := range array {
		if d < num && d > data {
			data = d
			addr = it
		}
	}
	return
}
