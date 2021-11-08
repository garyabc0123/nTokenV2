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

func ReadSearchQueue(path string) (string, error) {
	file, err := os.ReadFile(path)
	return string(file), err

}

func ReadData(path string) (Document, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var stringList []string
	for scanner.Scan() {
		stringList = append(stringList, scanner.Text()) //read line mode
	}

	var ret Document = make(Document, len(stringList))
	b1 := goroutines.NewBatch(runtime.NumCPU()*2, goroutines.WithBatchSize(len(stringList)))

	//var token [][]string = make([][]string, len(stringList))
	for id, _ := range stringList {
		idx := id

		b1.Queue(func() (interface{}, error) {

			token := strings.Split(stringList[idx], " ") //split strings by space
			var id int = 0
			var buf bytes.Buffer
			tempOutSen := make(Sentence, 0, len(token))
			for i, _ := range token {
				switch {
				case len(token[i]) > 1 && token[i][0] == '(' && token[i][len(token[i])-1] == ')':
					partofspeech, err := strconv.Atoi(token[i][1 : len(token[i])-1])
					if err != nil {
						return nil, err
					}
					tempOutSen = append(tempOutSen, &WordAndPartOfSpeechPair{Id: id, Word: []rune(buf.String()), PartOfSeech: partofspeech})
					ret[idx] = &tempOutSen
					id++
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
