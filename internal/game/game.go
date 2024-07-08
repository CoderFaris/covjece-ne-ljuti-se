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

var player1Path = [][]int{

	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0},
	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 10},
	{6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
}

var player2Path = [][]int{

	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 10},
	{6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 0},
}

var p1PathIndex int
var p2PathIndex int

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

// board for calculations

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

	p2out = true
	b[4][0] = "1"
	DrawCurrentBoard(b)

	return true

}

func P1Move(posX, posY, pathIndex int) (int, int, int) {
	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)

	newPathIndex := getNextPosition(player1Path, pathIndex, roll)
	newPosX := player1Path[newPathIndex][0]
	newPosY := player1Path[newPathIndex][1]

	if isValidPosition(newPosX, newPosY, b) {
		b[posX][posY] = "0"
		b[newPosX][newPosY] = "1"
		DrawCurrentBoard(b)
		return newPosX, newPosY, newPathIndex
	}

	fmt.Println("Invalid position")
	b[posX][posY] = "1"
	DrawCurrentBoard(b)
	return posX, posY, pathIndex
}

func P2Move(posX, posY, pathIndex int) (int, int, int) {
	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)

	newPathIndex := getNextPosition(player2Path, pathIndex, roll)
	newPosX := player2Path[newPathIndex][0]
	newPosY := player2Path[newPathIndex][1]

	if isValidPosition(newPosX, newPosY, b) {
		b[posX][posY] = "0"
		b[newPosX][newPosY] = "1"
		DrawCurrentBoard(b)
		return newPosX, newPosY, newPathIndex
	}

	fmt.Println("Invalid position")
	b[posX][posY] = "1"
	DrawCurrentBoard(b)
	return posX, posY, pathIndex
}

func getNextPosition(path [][]int, currentIndex int, roll int) int {

	newPosition := currentIndex + roll
	if newPosition >= len(path) {

		newPosition = newPosition % len(path)

	}

	return newPosition

}

func isValidPosition(posX, posY int, b Board) bool {
	if posX < 0 || posX >= len(b) || posY < 0 || posY >= len(b[0]) {
		return false
	}
	cell := b[posX][posY]
	return cell != "#" && cell != "h"
}

/*
	func getValidPositions(b Board) [][]int {
		var validPositions [][]int

		for i, row := range b {
			for j, cell := range row {

				if cell == "0" {
					validPositions = append(validPositions, []int{i, j})
				}
			}
		}

		return validPositions
	}

	func containsPosition(posX, posY int) bool {
		for _, position := range validPositions {
			if position[0] == posX && position[1] == posY {
				return true
			}
		}
		return false
	}
*/
func Game() {

	var currentPlayer int
	var choice int

	DrawCurrentBoard(b)

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
							p1PosX, p1PosY, p1PathIndex = P1Move(p1PosX, p1PosY, p1PathIndex)
							fmt.Println(p1PosX, p1PosY)

						}
					}
				} else {
					p1PosX, p1PosY, p1PathIndex = P1Move(p1PosX, p1PosY, p1PathIndex)
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
							p2PosX, p2PosY, p2PathIndex = P2Move(p2PosX, p2PosY, p2PathIndex)
							fmt.Println(p2PosX, p2PosY)
						}
					}
				} else {
					p2PosX, p2PosY, p2PathIndex = P2Move(p2PosX, p2PosY, p2PathIndex)
					fmt.Println(p2PosX, p2PosY)
				}
				currentPlayer = 0
			} else {
				fmt.Println("Skipping turn.")
				currentPlayer = 0
			}
		}
	}

}
