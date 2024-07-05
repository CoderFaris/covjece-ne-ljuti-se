package game

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/fatih/color"
)

type Board [11][11]string

var p1out, p2out bool

var p1PosX = 10
var p1PosY = 4

var p2PosX = 4
var p2PosY = 0

// initializing and returning board
func GetBoard() Board {
	red := color.New(color.FgRed).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	black := color.New(color.FgBlack).SprintFunc()

	board := Board{
		{"", "", "#", "#", "0", "0", "0", "#", "#", "", ""},
		{"", "", "#", "#", "0", "h", "0", "#", "#", "", ""},
		{"#", "#", "#", "#", "0", "h", "0", "#", "#", "#", "#"},
		{"#", "#", "#", "#", "0", "h", "0", "#", "#", "#", "#"},
		{"0", "0", "0", "0", "0", "h", "0", "0", "0", "0", "0"},
		{"0", "h", "h", "h", "h", "#", "h", "h", "h", "h", "0"},
		{"0", "0", "0", "0", "0", "h", "0", "0", "0", "0", "0"},
		{"#", "#", "#", "#", "0", "h", "0", "#", "#", "#", "#"},
		{"#", "#", "#", "#", "0", "h", "0", "#", "#", "#", "#"},
		{"", "", "#", "#", "0", "h", "0", "#", "#", "", ""},
		{"", "", "#", "#", "0", "0", "0", "#", "#", "", ""},
	}

	// out positions
	board[0][0] = red("0")
	board[0][1] = red("0")
	board[1][0] = red("0")
	board[1][1] = red("0")

	board[0][9] = blue("0")
	board[0][10] = blue("0")
	board[1][9] = blue("0")
	board[1][10] = blue("0")

	board[9][0] = yellow("0")
	board[9][1] = yellow("0")
	board[10][0] = yellow("0")
	board[10][1] = yellow("0")

	board[9][9] = black("0")
	board[9][10] = black("0")
	board[10][9] = black("0")
	board[10][10] = black("0")

	// home (ending) positions
	board[1][5] = blue("h")
	board[2][5] = blue("h")
	board[3][5] = blue("h")
	board[4][5] = blue("h")

	board[6][5] = yellow("h")
	board[7][5] = yellow("h")
	board[8][5] = yellow("h")
	board[9][5] = yellow("h")

	board[5][1] = red("h")
	board[5][2] = red("h")
	board[5][3] = red("h")
	board[5][4] = red("h")

	board[5][6] = black("h")
	board[5][7] = black("h")
	board[5][8] = black("h")
	board[5][9] = black("h")

	// starting positions
	board[10][4] = yellow("A")
	board[4][0] = red("A")
	board[0][6] = blue("A")
	board[6][10] = black("A")

	return board
}

// current board state
func DrawCurrentBoard(board Board) {
	for _, row := range board {
		for _, cell := range row {
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}

func RollDice() int {

	return rand.IntN(7-1) + 1

}

func wantsToRoll(choice int) bool {

	return choice == 1

}

func isOut(roll int) bool {

	return roll == 6

}

// board that the players will use
var b = GetBoard()

func P1Out() bool {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	if !isOut(roll) {

		fmt.Println("Roll must be a six")
		return false

	}

	// player managed to go to the starting point
	p1out = true
	b[10][4] = "1"
	DrawCurrentBoard(b)

	return true

}

func P2Out() bool {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	if !isOut(roll) {

		fmt.Println("Roll must be a six")
		return false
	}

	// player managed to go to the starting point
	p2out = true
	b[4][0] = "1"
	DrawCurrentBoard(b)

	return true

}

func P1Move(posX, posY int, b Board) (int, int) {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	b[posX][posY] = "0"
	b[posX-roll][posY] = "1"
	DrawCurrentBoard(b)
	return posX - roll, posY

}

func P2Move(posX, posY int, b Board) (int, int) {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	b[posX][posY] = "0"
	b[posX][posY+roll] = "1"
	DrawCurrentBoard(b)
	return posX, posY + roll

}

func calculateMove(posX, posY int, b Board) bool {

	return true

}

func Game() {

	var currentPlayer int
	var choice int

	for {
		if currentPlayer == 0 {
			fmt.Print("(P1) Roll dice? (1 for Yes, any key for No): ")
			fmt.Scan(&choice)
			if wantsToRoll(choice) {
				if !p1out {
					if P1Out() {
						fmt.Print("(P1) Roll again? (1 for Yes, any key for No): ")
						fmt.Scan(&choice)
						if wantsToRoll(choice) {
							b[10][4] = "0"
							p1PosX, p1PosY = P1Move(p1PosX, p1PosY, b)
							fmt.Println(p1PosX, p1PosY)

						}
					}
				} else {
					p1PosX, p1PosY = P1Move(p1PosX, p1PosY, b)
					fmt.Println(p1PosX, p1PosY)
				}
				currentPlayer = 1
			} else {
				fmt.Println("Skipping turn.")
				currentPlayer = 1
			}
		} else {
			fmt.Print("(P2) Roll dice? (1 for Yes, 0 for No): ")
			fmt.Scan(&choice)
			if wantsToRoll(choice) {
				if !p2out {
					if P2Out() {
						fmt.Print("(P2) Roll again? (1 for Yes, 0 for No): ")
						fmt.Scan(&choice)
						if wantsToRoll(choice) {
							b[4][0] = "0"
							p2PosX, p2PosY = P2Move(p2PosX, p2PosY, b)
							fmt.Println(p2PosX, p2PosY)
						}
					}
				} else {
					p2PosX, p2PosY = P2Move(p2PosX, p2PosY, b)
					fmt.Println(p1PosX, p1PosY)
				}
				currentPlayer = 0
			} else {
				fmt.Println("Skipping turn.")
				currentPlayer = 0
			}
		}
	}

}
