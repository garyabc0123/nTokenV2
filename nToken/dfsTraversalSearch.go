package nToken

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"reflect"
)

/*
                                         ------------ <- offset
                                         |          |
|___________________________________________________|
     ^                                              ^
     |                                              |
     base                                           Max
*/
type stackStatus struct {
	beginData Position
	endID     int
	nowID     int // -1: not found, -2: default
}

func NoName(position []DocumentPositionList, dist DistList) (retData [][]SentencePositionList) {
	//TODO give name

	retData = make([][]SentencePositionList, len(position[0]))
	for sentenceID := 0; sentenceID < len(position[0]); sentenceID++ {
		for startAdr := 0; startAdr < len(position[0][sentenceID]); startAdr++ {
			stack := arraystack.New()
			stack.Push(stackStatus{nowID: startAdr, beginData: -2, endID: startAdr})

			//begin loop dfs
			for {
				if stack.Empty() {
					break
				}
				nowLevel := stack.Size() - 1
				t, _ := stack.Peek()
				now := t.(stackStatus)

				if now.nowID == -1 {
					stack.Pop()
					if !stack.Empty() {
						temp, _ := stack.Peek()
						temp2 := temp.(stackStatus)
						temp2.nowID--
						stack.Pop()
						stack.Push(temp2)

					}
					continue
				}

				if now.nowID >= 0 && position[nowLevel][sentenceID][now.nowID] <= now.beginData {
					stack.Pop()
					if !stack.Empty() {
						temp, _ := stack.Peek()
						temp2 := temp.(stackStatus)
						temp2.nowID--
						stack.Pop()
						stack.Push(temp2)

					}
					continue
				}

				if nowLevel == len(dist) && now.nowID != -2 && now.endID != -2 {

					outTemp := make(SentencePositionList, 0)
					id := stack.Size() - 1
					for !stack.Empty() {
						t, _ := stack.Peek()
						stack.Pop()
						temp := t.(stackStatus)
						outTemp = append(outTemp, position[id][sentenceID][temp.nowID])
						id--

					}
					reverseSlice(outTemp)
					retData[sentenceID] = append(retData[sentenceID], outTemp)
					break
				}

				if now.nowID == -2 && now.endID == -2 {
					stack.Pop()
					temp, _ := findingFloor(position[nowLevel][sentenceID], now.beginData+Position(dist[nowLevel-1]))
					now.endID = temp
					now.nowID = temp

					stack.Push(now)
					continue
				}

				nextEnd, _ := findingFloor(position[nowLevel+1][sentenceID], position[nowLevel][sentenceID][now.nowID]+Position(dist[nowLevel]))
				if nextEnd == -1 {
					now.nowID--
					stack.Pop()
					stack.Push(now)
					continue
				} else {
					stack.Push(stackStatus{position[nowLevel][sentenceID][now.nowID], -2, -2})
					continue
				}

			}
		}

	}
	return
}

//不超過number的list 中最大的一個
//return position, data
func findingFloor(list []Position, number Position) (int, Position) {
	if number < list[0] {
		return -1, -1
	}

	if number >= list[len(list)-1] {
		return len(list) - 1, list[len(list)-1]
	}

	var begin int = 0
	var end int = len(list) - 1
	var mid int
	for end > begin {
		mid = (end-begin)/2 + begin
		if list[mid] <= number && list[mid+1] > number {
			return mid, list[mid]
		} else if list[mid+1] < number {
			begin = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1, -1
}

//不小於number的list中最小的一個
func findingCeiling(list []Position, number Position) (int, Position) {
	if number > list[len(list)-1] {
		return -1, -1
	}
	if number < list[0] {
		return 0, list[0]
	}
	var begin int = 0
	var end int = len(list)
	var mid int
	for end > begin {
		mid = (end-begin)/2 + begin
		if list[mid] >= number && list[mid+1] < number {
			return mid, list[mid]
		} else if list[mid+1] < number {
			begin = mid + 1
		} else {
			end = mid - 1
		}

	}
	return -1, -1

}

func reverseSlice(s interface{}) {
	size := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, size-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}
