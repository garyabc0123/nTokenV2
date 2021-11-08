package nToken

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"os"
	"testing"
)

type testData struct {
	searchKey string
	word      WordAndPartOfSpeechPair
	ans       bool
}

func TestTreetraversal(t *testing.T) {
	var c []*ParseTree
	var err error
	var ans bool

	file, err := os.Open("position_test.json")
	if err != nil {
		t.Fatal(err)
	}
	testjson, err := simplejson.NewFromReader(file)
	if err != nil {
		t.Fatal(err)
	}

	for i, row := range testjson.MustArray() {
		c, _ = Compiler(row.(map[string]interface{})["searchKey"].(string))
		tempId, _ := row.(map[string]interface{})["word"].(map[string]interface{})["Id"].(json.Number).Int64()
		tempPoS, _ := row.(map[string]interface{})["word"].(map[string]interface{})["PartOfSpeech"].(json.Number).Int64()
		ans, err = Treetraversal(c[0], WordAndPartOfSpeechPair{
			Id:           int(tempId),
			Word:         []rune(row.(map[string]interface{})["word"].(map[string]interface{})["Word"].(string)),
			PartOfSpeech: int(tempPoS),
		})
		if err != nil {
			t.Error(err)
		}
		if ans != row.(map[string]interface{})["ans"].(bool) {
			t.Error(i)
		} else {
			println(i, " was passed! ")
		}
	}

}
