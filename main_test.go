package main

import (
	"testing"
)

func TestMakeBoggleIndex(t *testing.T) {
	boggleBoard := [][]rune{
		{'A', 'B'},
		{'C', 'D'},
	}

	idx := NewBoggleIndex(boggleBoard)

	_, existsA00 := idx['A'][0][0]
	_, existsA01 := idx['A'][0][1]
	_, existsB00 := idx['B'][0][0]
	_, existsB01 := idx['B'][0][1]
	_, existsC10 := idx['C'][1][0]
	_, existsC00 := idx['C'][0][0]
	_, existsD11 := idx['D'][1][1]
	if !existsA00 {
		t.Errorf("Expected A at 00.")
	}
	if existsA01 {
		t.Errorf("Expected A not at 01.")
	}
	if existsB00 {
		t.Errorf("Expected B not at 00.")
	}
	if !existsB01 {
		t.Errorf("Expected B at 01.")
	}
	if !existsC10 {
		t.Errorf("Expected C at 10.")
	}
	if existsC00 {
		t.Errorf("Expected C not at 00.")
	}
	if !existsD11 {
		t.Errorf("Expected D at 11.")
	}
}

func TestBoggleBoard(t *testing.T) {
	var boggleBoard = BoggleBoard{
		{'G', 'I', 'Z'},
		{'U', 'E', 'K'},
		{'Q', 'S', 'E'},
	}

	if !boggleBoard.CheckWordInBoard("GEEK") {
		t.Errorf("Excepected to find GEEK")
	}
	if boggleBoard.CheckWordInBoard("FOR") {
		t.Errorf("Excepected not to find FOR")
	}
	if !boggleBoard.CheckWordInBoard("QUIZ") {
		t.Errorf("Excepected to find QUIZ")
	}
	if boggleBoard.CheckWordInBoard("GO") {
		t.Errorf("Excepected not to find GO")
	}
}
