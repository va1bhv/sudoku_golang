package main

import "fmt"

func drawBoard(board [9][9]int) {
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
func valid(num int, pos [2]int, board [9][9]int) bool { // Position will be (Row, Column)
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
func findEmpty(board [9][9]int) [2]int{
	for i:=0; i<9; i++ {
		for j := 0; j<9; j++ {
			if board[i][j] == 0 {
				return [2]int{i, j}
			}
		}
	}
	return [2]int{10,10}
}

func main(){
	st := [9][9]int{
		[9]int{1,2,3,0,0,0,0,0,0},
		[9]int{4,5,6,0,0,0,0,0,0},
		[9]int{7,8,9,0,0,0,0,0,0},
		[9]int{0,0,0,1,2,3,0,0,0},
		[9]int{0,0,0,4,5,6,0,0,0},
		[9]int{0,0,0,7,8,9,0,0,0},
		[9]int{0,0,0,0,0,0,1,2,3},
		[9]int{0,0,0,0,0,0,4,5,6},
		[9]int{0,0,0,0,0,0,7,8,9},
	}
	// drawBoard(st)

	fmt.Println(valid(5, [2]int{0, 3}, st))
	fmt.Println(findEmpty(st))
}
