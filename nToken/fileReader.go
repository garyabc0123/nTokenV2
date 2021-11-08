package nToken

import (
	"bufio"
	"bytes"
	"github.com/viney-shih/goroutines"
	"os"
	"runtime"
	"strconv"
	"strings"
)

//read seaech queue from file
func ReadSearchQueue(path string) (string, error) {
	file, err := os.ReadFile(path)
	return string(file), err

}

//read data from file and convert to Document
func ReadData(path string) (Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stringList []string //readline()
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text()) //read line mode
	}

	var ret Document = make(Document, len(stringList))                                       //return data
	b1 := goroutines.NewBatch(runtime.NumCPU()*2, goroutines.WithBatchSize(len(stringList))) //work pool
	for id, _ := range stringList {
		idx := id //backup, because id is a variable with is shared by multi-thread.

		b1.Queue(func() (interface{}, error) { //push into work queue

			token := strings.Split(stringList[idx], " ") //split strings by space
			var idInside int = 0
			var buf bytes.Buffer
			tempOutSen := make(Sentence, 0, len(token))
			for i, _ := range token {
				switch {
				case len(token[i]) > 1 && token[i][0] == '(' && token[i][len(token[i])-1] == ')':
					partofspeech, err := strconv.Atoi(token[i][1 : len(token[i])-1])
					if err != nil {
						return nil, err
					}
					tempOutSen = append(tempOutSen, &WordAndPartOfSpeechPair{Id: idInside, Word: []rune(buf.String()), PartOfSpeech: partofspeech})
					ret[idx] = &tempOutSen
					idInside++
					buf.Reset()
				default:
					buf.WriteString(token[i])

				}

			}

			return nil, nil
		})
	}
	b1.QueueComplete()
	b1.Close()

	return ret, nil

}
