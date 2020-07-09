package game

import (
	"fmt"
)

//Move is a tuple with a bet and board position
type Move [2]int

//State holds the current state of the game
type State struct {
	player1Points int
	player2Points int
	board         []int
	history       []State
}

//Step updates the environment after receiving moves from both players
func (game State) Step(p1, p2 Move) State {
	if game.board[p1[1]] != 0 || game.board[p2[1]] != 0 {
		fmt.Println("Invalid Move")
		return game
	}
	if p1[0] > game.player1Points || p2[0] > game.player2Points {
		fmt.Println("Invalid Bet")
		return game
	}
	if p1[0] > p2[0] {

		game.board[p1[1]] = 1
	}
	return
}
