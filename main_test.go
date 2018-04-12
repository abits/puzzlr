package main

import (
	"fmt"
	"testing"
	"reflect"
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

type testListBoard struct {
	boardList []Board
	board Board
	expect bool
}

type testListBoardList struct {
	boardList []Board
	board Board
	reducedList []Board
}

func dump(boardList []Board) {
	for _, b := range(boardList) {
		b.print()		
	}
}

var testBoards = []Board{
	Board{{0, 5, 6},{2, 3, 1},{7, 4, 8}},
	Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}},
	Board{{0, 5, 9},{2, 3, 1},{7, 4, 8}},
	Board{{0, 5, 9},{2, 1},{7, 4, 8}},
	Board{{1, 2, 3},{8, 0, 4},{7, 6, 5}},
	Board{{5, 2, 3},{8, 0, 4},{7, 6, 1}},
}

var reducedBoards = []Board{
	Board{{0, 5, 6},{2, 3, 1},{7, 4, 8}},
	Board{{0, 5, 9},{2, 3, 1},{7, 4, 8}},
	Board{{0, 5, 9},{2, 1},{7, 4, 8}},
	Board{{1, 2, 3},{8, 0, 4},{7, 6, 5}},
	Board{{5, 2, 3},{8, 0, 4},{7, 6, 1}},
}

var delBoards = []Board{
	Board{{0, 5, 6},{2, 3, 1},{7, 4, 8}},
	Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}},
	Board{{0, 5, 9},{2, 3, 1},{7, 4, 8}},
	Board{{0, 5, 9},{2, 1},{7, 4, 8}},
	Board{{1, 2, 3},{8, 0, 4},{7, 6, 5}},
	Board{{5, 2, 3},{8, 0, 4},{7, 6, 1}},
}

var delTests = []testListBoardList{
	{delBoards, Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}}, reducedBoards},
}

func TestDel(t *testing.T) {
	for _, pair := range(delTests) {
		expectedList := del(pair.boardList, pair.board)
		if contains(expectedList, pair.board) {
			t.Error(
                "For", pair.boardList,
                "expected", reducedBoards,
                "got", expectedList,
             )
		}

	}

}

var validationTests = []testPair{
	{ testBoards[0], true },
	{ testBoards[1], true },
	{ testBoards[2], false },
	{ testBoards[3], false },
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

var diffTests = []testTripleBoards{
	{ testBoards[5], testBoards[4], 8 },
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
}

var findPosTests = []testTripleInt{
	{ testBoards[0], 1, 1, 2},
	{ testBoards[0], 8, 2, 2},
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

var containsTests = []testListBoard {
	{ testBoards, testBoards[0], true },
	{ testBoards, testBoards[1], true },
	{ testBoards, testBoards[2], true },
	{ testBoards, testBoards[3], true },
	{ testBoards, Board{{0, 5, 9},{2, 3, 2},{7, 4, 8}}, false },
	{ testBoards, testBoards[5], true },
}

func TestContains(t *testing.T) {
	for _, pair := range(containsTests) {
		gotBool := contains(pair.boardList, pair.board)
		if gotBool != pair.expect {
			t.Error(
        		"For", fmt.Sprintf("%v and %v", pair.boardList, pair.board),
        		"expected", fmt.Sprintf("%v", pair.expect),
        		"got", fmt.Sprintf("%v", gotBool),

     		)
		}
	}
}

type testMoveTiles struct {
	orig Board
	tile []int
	expect Board
}

var moveTileTests = []testMoveTiles{
	{Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}}, []int{1,0}, Board{{3, 1, 2},{0, 4, 5},{6, 7, 8}}},
	{Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}}, []int{0,1}, Board{{1, 0, 2},{3, 4, 5},{6, 7, 8}}},
	{Board{{4, 1, 2},{3, 0, 5},{6, 7, 8}}, []int{0,1}, Board{{4, 0, 2},{3, 1, 5},{6, 7, 8}}},
	{Board{{4, 1, 2},{3, 0, 5},{6, 7, 8}}, []int{2,1}, Board{{4, 1, 2},{3, 7, 5},{6, 0, 8}}},
	{Board{{4, 1, 2},{3, 0, 5},{6, 7, 8}}, []int{1,2}, Board{{4, 1, 2},{3, 5, 0},{6, 7, 8}}},
}

func TestMoveTile(t *testing.T) {
	for _, pair := range(moveTileTests) {
		pair.orig.moveTile(pair.tile)
		if !(pair.expect.equal(pair.orig)) {
			t.Error(
        		"For", fmt.Sprintf("%v and %v", pair.orig, pair.tile),
        		"expected", fmt.Sprintf("%v", pair.expect),
        		"got", fmt.Sprintf("%v", pair.orig),
			)
		}
	}
}

type testFindNeighbors struct {
	board Board
	expect [][]int
}
var findNeighborsTest = []testFindNeighbors{
	{Board{{0, 1, 2},{3, 4, 5},{6, 7, 8}}, [][]int{{1, 0}, {0, 1}}},
	{Board{{1, 2, 3},{4, 0, 5},{6, 7, 8}}, [][]int{{2, 1}, {0, 1}, {1, 0}, {1, 2}}},
	{Board{{1, 2, 3},{4, 5, 6},{7, 0, 8}}, [][]int{{1, 1}, {2, 0}, {2, 2}}},
}

func TestFindNeighbors(t *testing.T) {
	for _, pair := range(findNeighborsTest) {
		result := pair.board.findNeighbors()
		if !(reflect.DeepEqual(result, pair.expect)) {
			t.Error(
        		"For", fmt.Sprintf("%v", pair.board),
        		"expected", fmt.Sprintf("%v", pair.expect),
        		"got", fmt.Sprintf("%v", result),
			)
		}
	}
} 
