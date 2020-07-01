package main

import "fmt"

func drawBoard(board [][]int) {
	for i, row := range board{
		if i%3 == 0 && i != 0{
			fmt.Print("--- --- ---")
			fmt.Println()
		}
		for j, val := range row{
			if j%3 == 0 && j != 0{
				fmt.Print("|")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

func valid(num int, pos []int, board [][]int) bool { // Position will be (Row, Column)
	// Check row
	for i:=0; i<9; i++ { 
		if num == board[i][pos[0]] && i != pos[0] {
			return false
		}
	}
	// Check column
	for i:=0; i<9; i++ { 
		if num == board[pos[1]][i] && i != pos[1] {
			return false
		}
	}
	// Check small square
	for i:=pos[0]; i<pos[0] + 3; i++ {
		for j:=pos[1]; j<pos[1] + 3; j++ {
			if board[i][j] == num {
				return false
			}
		}
	}
	return true
}

func findEmpty(board [][]int) []int{
	for i:=0; i<9; i++ {
		for j := 0; j<9; j++ {
			if board[i][j] == 0 {
				return []int{i, j}
			}
		}
	}
	return nil
}

func solve(board [][]int) bool {
	empty := findEmpty(board)
	if empty == nil {
		return true
	}
	row := empty[0]
	col := empty[1]
	for i:=1;i<10;i++ {
		if valid(i, empty, board) {
			board[row][col] = i
			if solve(board) {
				return true
			}
			col[board][row] = 0
		}
	}
	return false
}

func main(){
	st := [][]int{
		[]int{1,2,3,0,0,0,0,0,0},
		[]int{4,5,6,0,0,0,0,0,0},
		[]int{7,8,9,0,0,0,0,0,0},
		[]int{0,0,0,1,2,3,0,0,0},
		[]int{0,0,0,4,5,6,0,0,0},
		[]int{0,0,0,7,8,9,0,0,0},
		[]int{0,0,0,0,0,0,1,2,3},
		[]int{0,0,0,0,0,0,4,5,6},
		[]int{0,0,0,0,0,0,7,8,9},
	}
	// drawBoard(st)

	drawBoard(st)
	fmt.Println("* * * * * * * * *")
	a := solve(st)
	drawBoard(st)
	fmt.Println(a)
}
