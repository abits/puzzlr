package main

// 1. Create two lists, L and L seen. At the beginning, 
//    L contains only the initial state, and L seen is empty.
// 2. Let n be the first element of L. Compare this state with the final state. 
//    If they are identical, stop with success.
// 3. Apply to n all available search operators, thus obtaining a set of
//    new states. Discard those states that already exist in L seen . As for
//    the rest, sort them by the evaluation function and place them 
//    at the front of L.
// 4. Transfer n from L into the list, L seen , of the states that
//    have been investigated.
// 5. If L = ∅, stop and report failure. Otherwise, go to 2.
// Excerpt from: Miroslav Kubat. "An Introduction to Machine Learning."

import (
    "fmt"
	"sort"
	"reflect"
)

type Board [][]int
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

func main() {
	//board1.print()
	//fmt.Println(board1.flatten())
	//fmt.Printf("%v", board())
}
