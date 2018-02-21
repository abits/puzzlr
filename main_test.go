package main

import (
	"fmt"
	"testing"
)

type testPair struct{
	board Board
	expect bool
}

type testTriple struct{
	board1 Board
	board2 Board
	expect bool
}

type testTripleInt struct {
	board Board
	lookfor int
	row int
	num int
}

type testTripleBoards struct {
	board1 Board
	board2 Board
	expect int
}

var testBoards = []Board{
	Board{{0, 5, 6},{2, 3, 1},{7, 4, 8}},
	Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}},
	Board{{0, 5, 9},{2, 3, 1},{7, 4, 8}},
	Board{{0, 5, 9},{2, 1},{7, 4, 8}},
	Board{{1, 2, 3},{8, 0, 4},{7, 6, 5}},
	Board{{5, 2, 3},{8, 0, 4},{7, 6, 1}},
}

var validationTests = []testPair{
	{ testBoards[0], true },
	{ testBoards[1], true },
	{ testBoards[2], false },
	{ testBoards[3], false },
}

var equalTests = []testTriple{
	{ testBoards[0], testBoards[0], true },
	{ testBoards[1], testBoards[1], false },
	{ testBoards[1], testBoards[0], false },
	{ testBoards[3], testBoards[2], false },
}

var findPosTests = []testTripleInt{
	{ testBoards[0], 1, 1, 2},
	{ testBoards[0], 8, 2, 2},
}

var diffTests = []testTripleBoards{
	{ testBoards[5], testBoards[4], 8 },
}

func TestValidate(t *testing.T) {
	for _, pair := range(validationTests) {
		if pair.board.validate() != pair.expect {
			t.Error(
        		"For", pair.board,
        		"expected", pair.expect,
        		"got", pair.board.validate(),
     		)
		} 
	}	
}

func TestDiff(t *testing.T) {
	for _, pair := range(diffTests) {
		if pair.board1.diff(pair.board2) != pair.expect {
			t.Error(
        		"For", pair.board1,
        		"expected", pair.expect,
        		"got", pair.board1.diff(pair.board2),
     		)
		} 
	}	
	t.Logf("%d", testBoards[5].diff(testBoards[4]))

}

func TestFindPos(t *testing.T) {
	for _, pair := range(findPosTests) {
		gotRow, gotNum := pair.board.findPos(pair.lookfor)
		if (gotRow != pair.row) || (gotNum != pair.num) {
			t.Error(
        		"For", fmt.Sprintf("%v and %d", pair.board, pair.lookfor),
        		"expected", fmt.Sprintf("%d %d", pair.row, pair.num),
        		"got", fmt.Sprintf("%d %d", gotRow, gotNum),

     		)
		} 
	}	
}

