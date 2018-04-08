package main

// 1. Create two lists, L and L seen. At the beginning, 
//    L contains only the initial state, and L seen is empty.
// 2. Let n be the first element of L. Compare this state with the final state. 
//    If they are identical, stop with success.
// 3. Apply to n all available search operators, thus obtaining a set of
//    new states. Discard those states that already exist in L seen. As for
//    the rest, sort them by the evaluation function and place them 
//    at the front of L.
// 4. Transfer n from L into the list, L seen, of the states that
//    have been investigated.
// 5. If L = ∅, stop and report failure. Otherwise, go to 2.
// Excerpt from: Miroslav Kubat. "An Introduction to Machine Learning."

import (
    "fmt"
	"sort"
	"reflect"
)

type Board [3][3]int
type Boards []Board
type BoardsSeen []Board

func (board Board) print() {
	fmt.Printf("+---+---+---+\n")
	for _, r := range (board) {
		for _, i := range(r) {
			if i != 0 {
				fmt.Printf("| %d ", i)
			} else {
				fmt.Printf("|   ")
			}	
		}
		fmt.Printf("|\n")
	}
	fmt.Printf("+---+---+---+\n")
}

func (board1 Board) equal(board2 Board) bool {
	return reflect.DeepEqual(board1, board2)
}

func (board Board) findNeighbors() (neighbors [][]int){
	zeroRow, zeroCol := board.findPos(0)
	if (zeroRow + 1) <= 2 {
		neighbors = append(neighbors, []int{zeroRow + 1, zeroCol})
	}
	if (zeroRow -1  >= 0) {
		neighbors = append(neighbors, []int{zeroRow - 1, zeroCol})
	}
	if (zeroCol -1 >= 0 ) {
		neighbors = append(neighbors, []int{zeroRow, zeroCol - 1})
	}
	if (zeroCol + 1 <= 2) {
		neighbors = append(neighbors, []int{zeroRow, zeroCol + 1})
	}
	return
}

func (board *Board) moveTile(pos []int) {
	row, num := board.findPos(0)
	board[row][num] = board[pos[0]][pos[1]]
	board[pos[0]][pos[1]] = 0
}

func (orig Board) search() (newStates Boards) {
	neighbors := orig.findNeighbors()
	for _, pos := range(neighbors) {
		board := orig
		board.moveTile(pos)
		newStates = append(newStates, board)
	}
	return
}

func (board1 Board) diff(board2 Board) (delta int) {
	for i := 1; i < 10; i++ {
		a, b := board1.findPos(i)
		c, d := board2.findPos(i)
		delta = delta + abs(a-c) + abs(b-d)
	}
	return 
}

func abs(n int) int {
	if n < 0 {
   		n = -n
	}
	return n
}

//func deepCopy(dst Board, src Board) { 
	//for i := range src {
        //dst[i] = make([3]int, len(src[i]))
        //copy(dst[i], src[i])
	//}
//}

func (board Board) validate() bool {
	if len(board) != 3 {
		return false
	}
	flat := board.flatten()
	if len(flat) != 9 {
		return false
	}
	sort.Ints(flat)
	shouldBe := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	for i, j := range(flat) {
		if j != shouldBe[i] {
			return false
		}
	}
	return true
}

// find position of value in slice
func (board Board) findPos(value int) (i int, j int) {
	for i, row := range(board) {
		for j, num := range(row) {
			if value == num {
				return i, j
			}
		}
	}
	return
}

func (board Board) flatten() (flat []int) {
	for _, r := range (board) {
		for _, i := range(r) {
			flat = append(flat, i)
		}
	}
	return
}

func contains(boards []Board, board Board) bool {
	for _, val := range boards {
		if val == board {
			return true 
		}
	}
	return false
}

func del(boards []Board, board Board) ([]Board) {
	for i, val := range boards {
		if val == board {
			boards = append(boards[:i], boards[i+1:]...)
		}
	}
	return boards
}

func removeSeen(states []Board, boardsSeen BoardsSeen) (unseen []Board) {
	for _, i := range states {
		if contains(boardsSeen, i) {
			states = del(states, i)
		}
	}
	unseen = states
	return
}

func sortStates(states []Board, goal Board) (sorted []Board) {
	
	m := make(map[int]Board)
	
	for _, i := range states {
		m[i.diff(goal)] = i
	}
	// sort remaining states by their distance to goal
	var keys []int
	for k := range m {
        keys = append(keys, k)
    }
    sort.Ints(keys)
	for _, k := range keys {
    	sorted = append(sorted, m[k])
    }
	
	return
}

func debug(prefix string, states []Board, goal Board) {
	fmt.Println(prefix)
	for _, i := range states {
		i.print()
		fmt.Printf("%v\n", i.diff(goal))
	}
}

func main() {
	initial := Board{{7, 5, 6},{2, 3, 1},{0, 4, 8}}
	goal := Board{{1, 2, 3}, {4, 0, 5}, {6, 7, 8}}
	boards := Boards{initial}
	boardsSeen := BoardsSeen{}
	firstBoard := boards[0]
	for !(goal.equal(firstBoard)) {
		firstBoard.print()
		newStates := firstBoard.search()
		newStates = removeSeen(newStates, boardsSeen)
		newStates = sortStates(newStates, goal)
		boardsSeen = append(boardsSeen, boards[0])
		boards = boards[1:]
		boards = append(newStates, boards...)
		firstBoard = boards[0]
	} 	
	firstBoard.print()
}
