package nToken

import "testing"

/*
| senID | 詞組0位置 | 詞組1位置 | 詞組2位置 | 詞組3位置 | 詞組4位置 | Answer            |
| ----- | --------- | --------- | --------- | --------- | --------- | ----------------- |
| 0     | 1         | 2         | 3         | 4         | 5         | True,1,2,3,4,5    |
| 1     | 1         | 13        | 14        | 15        | 16        | False,nil         |
| 2     | 1         | 5,7       | 6         | 8         | 10        | True,1,5,6,8,10   |
| 3     | 1         | 5,7       | 8,9       | 7         | 12        | False             |
| 4     | 1         | 5,10      | 8,9       | 11        | 12        | True,1,5,9,10,12  |
| 5     | 1         | 1         | 1         | 1         | 1         | False             |
| 6     | 1         | 5,10      | 8,11      | 12,14     | 15        | True,1,10,11,14,15|
| 7     | 1         |2          | 1         | 2         | 1         | False             |
| 8     |           |           |           |           |           |                   |
*/
func TestDfsRecursive(t *testing.T) {
	var position []DocumentPositionList = []DocumentPositionList{
		{ //word 0
			{ //sentence 0
				1,
			},
			{ //sentence 1
				1,
			},
			{ //sentence 2
				1,
			},
			{ //sentence 3
				1,
			},
			{ //sentence 4
				1,
			},
			{ //sentence 5
				1,
			},
			{ //sentence 6
				1,
			},
			{ //sentence 7
				1,
			},
		},
		{ //word 1
			{ //sentence 0
				2,
			},
			{ //sentence 1
				13,
			},
			{ //sentence 2
				5, 7,
			},
			{ //sentence 3
				5, 7,
			},
			{ //sentence 4
				5, 10,
			},
			{ //sentence 5
				1,
			},
			{ //sentence 6
				5, 10,
			},
			{ //sentence 7
				2,
			},
		},
		{ //word 2
			{ //sentence 0
				3,
			},
			{ //sentence 1
				14,
			},
			{ //sentence 2
				6,
			},
			{ //sentence 3
				8, 9,
			},
			{ //sentence 4
				8, 9,
			},
			{ //sentence 5
				1,
			},
			{ //sentence 6
				8, 11,
			},
			{ //sentence 7
				1,
			},
		},
		{ //word 3
			{ //sentence 0
				4,
			},
			{ //sentence 1
				15,
			},
			{ //sentence 2
				8,
			},
			{ //sentence 3
				7,
			},
			{ //sentence 4
				11,
			},
			{ //sentence 5
				1,
			},
			{ //sentence 6
				12, 14,
			},
			{ //sentence 7
				2,
			},
		},
		{ //word 4
			{ //sentence 0
				5,
			},
			{ //sentence 1
				16,
			},
			{ //sentence 2
				10,
			},
			{ //sentence 3
				12,
			},
			{ //sentence 4
				12,
			},
			{ //sentence 5
				1,
			},
			{ //sentence 6
				15,
			},
			{ //sentence 7
				1,
			},
		},
	}
	var dist DistList = DistList{10, 10, 10, 10}
	var ans []bool = []bool{true, false, true, false, true, false, true, false}
	for i := 0; i < len(position[0]); i++ {
		x, q := dfsRecursive(&position, &dist, 0, PositionArray{TokenID: []Position{position[0][0][0]}, SentenceID: i})
		if x != ans[i] {
			t.Fail()
			println("Test ", i, "was fail. ")
		} else {
			println("Test ", i, "was success. ")
		}
		if q != nil {

			for a := 0; a < len(q); a++ {
				print(q[a], " ")
			}
			println()
		}
	}

}
