package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Dictionary []string
type BoggleBoard [][]rune
type BoggleIndex map[rune]map[int]map[int]bool

func NewBoggleIndex(boggleBoard BoggleBoard) BoggleIndex {
	index := map[rune]map[int]map[int]bool{}
	for i, row := range boggleBoard {
		for j, letter := range row {
			if index[letter] == nil {
				index[letter] = map[int]map[int]bool{
					i: map[int]bool{j: true},
				}
			} else if index[letter][i] == nil {
				index[letter][i] = map[int]bool{
					j: true,
				}
			} else {
				index[letter][i][j] = true
			}
		}
	}
	return index
}

func (index BoggleIndex) CheckWordInIndex(wordSlice []rune, x, y int) bool {
	_, exists := index[wordSlice[0]][x][y]
	if !exists {
		return false
	}
	if len(wordSlice) == 1 {
		return true
	}
	nextWordSlice := wordSlice[1:]

	checkLeft := index.CheckWordInIndex(nextWordSlice, x-1, y)
	if checkLeft {
		return true
	}

	checkRight := index.CheckWordInIndex(nextWordSlice, x+1, y)
	if checkRight {
		return true
	}

	checkTop := index.CheckWordInIndex(nextWordSlice, x, y-1)
	if checkTop {
		return true
	}

	checkBottom := index.CheckWordInIndex(nextWordSlice, x, y+1)
	if checkBottom {
		return true
	}

	checkTopLeft := index.CheckWordInIndex(nextWordSlice, x-1, y-1)
	if checkTopLeft {
		return true
	}

	checkBottomLeft := index.CheckWordInIndex(nextWordSlice, x-1, y+1)
	if checkBottomLeft {
		return true
	}

	checkTopRight := index.CheckWordInIndex(nextWordSlice, x+1, y-1)
	if checkTopRight {
		return true
	}

	checkBottomRight := index.CheckWordInIndex(nextWordSlice, x+1, y+1)
	if checkBottomRight {
		return true
	}

	return false
}

func (board BoggleBoard) CheckWordInBoard(word string) bool {
	index := NewBoggleIndex(board)
	wordSlice := []rune(word)
	letterIndex := index[wordSlice[0]]
	for i, mapj := range letterIndex {
		for j := range mapj {
			wordExists := index.CheckWordInIndex(wordSlice, i, j)
			if wordExists {
				return true
			}
		}
	}

	return false
}

func readInt(reader *bufio.Reader, delim byte) int {
	intString, _ := reader.ReadString(delim)
	parsedInt, _ := strconv.Atoi(intString[:len(intString)-1])
	return parsedInt
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	numTestCases := readInt(reader, '\n')

	var dict Dictionary
	for i := 0; i < numTestCases; i++ {
		numWordsInDict := readInt(reader, '\n')
		dict = Dictionary{}
		for j := 0; j < numWordsInDict-1; j++ {
			word, _ := reader.ReadString(' ')
			word = word[:len(word)-1]
			dict = append(dict, word)
		}
		word, _ := reader.ReadString('\n')
		word = word[:len(word)-1]
		dict = append(dict, word)
		numRows := readInt(reader, ' ')
		numCols := readInt(reader, '\n')
		boggleBoard := make(BoggleBoard, numRows)
		for col := 0; col < numCols; col++ {
			for row := 0; row < numRows; row++ {
				if col == 0 {
					boggleBoard[row] = make([]rune, numCols)
				}
				var letter string
				if col == numCols-1 && row == numRows-1 {
					letter, _ = reader.ReadString('\n')
				} else {
					letter, _ = reader.ReadString(' ')
				}
				boggleBoard[col][row] = rune(letter[0])
			}
		}

		for i, word := range dict {
			if boggleBoard.CheckWordInBoard(word) {
				if len(dict)-1 == i {
					fmt.Println(word)
				} else {
					fmt.Printf("%s ", word)
				}
			}
		}
	}
}
