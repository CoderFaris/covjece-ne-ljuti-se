package game

import (
	"fmt"
	"math/rand/v2"
	"time"

	"github.com/fatih/color"
)

// player path positions = 39 (0-39)
// laps needed for win = 4
// first to get 4 laps wins

type Board [11][11]string

type Player struct {
	playerOut      bool
	posX, posY     int
	totalLaps      int
	pathIndex      int
	path           [][]int
	startX, startY int
	laps           int
	color          string
}

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

var red = color.New(color.FgRed).SprintFunc()
var blue = color.New(color.FgBlue).SprintFunc()
var yellow = color.New(color.FgYellow).SprintFunc()
var black = color.New(color.FgBlack).SprintFunc()

var player1 = Player{
	posX:      10,
	posY:      4,
	path:      player1Path,
	startX:    10,
	startY:    4,
	playerOut: false,
	color:     yellow("1"),
}

var player2 = Player{
	posX:      4,
	posY:      0,
	path:      player2Path,
	startX:    4,
	startY:    0,
	playerOut: false,
	color:     red("1"),
}

func GetBoard() Board {

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

func checkWin(totalLaps int, player string) (bool, string) {

	if totalLaps == 4 && (player == "P1" || player == "P2") {

		return true, player

	}

	return false, ""

}

// board that the players will use
var b = GetBoard()

func playerOut(player *Player) bool {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	if !isOut(roll) {

		fmt.Println("Roll must be a six")
		return false

	}

	// player managed to go to the starting point
	player.playerOut = true
	b[player.startX][player.startY] = player.color
	DrawCurrentBoard(b)

	return true

}

func playerMove(player *Player) (int, int, int) {
	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)

	newIndex, laps := getNextPosition(player.path, player.pathIndex, roll)
	newPosX := player.path[newIndex][0]
	newPosY := player.path[newIndex][1]

	if isValidPosition(newPosX, newPosY, b) {
		b[player.posX][player.posY] = "0"
		b[newPosX][newPosY] = player.color
		DrawCurrentBoard(b)
		player.posX, player.posY, player.pathIndex = newPosX, newPosY, newIndex
		return newPosX, newPosY, laps
	} else {
		fmt.Println("Invalid position")
		b[player.posX][player.posY] = player.color
		DrawCurrentBoard(b)
		return player.posX, player.posY, 0
	}
}

func getNextPosition(path [][]int, currentIndex int, roll int) (int, int) {

	newPosition := currentIndex + roll
	laps := newPosition / len(path)
	newPosition = newPosition % len(path)

	return newPosition, laps

}

func isValidPosition(posX, posY int, b Board) bool {
	if posX < 0 || posX >= len(b) || posY < 0 || posY >= len(b[0]) {
		return false
	}
	cell := b[posX][posY]
	return cell != "#" && cell != "h"
}

func lapAction(lap int, player *Player) {

	switch lap {

	case 1:

		if player == &player1 {

			b[9][5] = "1"

		}

		if player == &player2 {

			b[5][1] = "1"

		}

	case 2:

		if player == &player1 {

			b[8][5] = "1"

		}

		if player == &player2 {

			b[5][2] = "1"

		}

	case 3:

		if player == &player1 {

			b[7][5] = "1"

		}

		if player == &player2 {

			b[5][3] = "1"

		}

	case 4:

		if player == &player1 {

			b[6][5] = "1"

		}

		if player == &player2 {

			b[5][4] = "1"

		}

	}

}

func getPlayerName(player *Player) string {

	if player == &player1 {

		return "P1"

	}

	return "P2"

}

func Game() {
	var currentPlayer *Player = &player1
	var otherPlayer *Player = &player2
	var choice int
	var gameEnded = false
	var winner string

	DrawCurrentBoard(b)

	for !gameEnded {
		fmt.Printf("(%v) Roll dice? (1 for Yes, any key for No): ", getPlayerName(currentPlayer))
		fmt.Scan(&choice)
		if wantsToRoll(choice) {
			if !currentPlayer.playerOut {
				if playerOut(currentPlayer) {
					fmt.Printf("(%v) Roll again? (1 for Yes, any key for No): ", getPlayerName(currentPlayer))
					fmt.Scan(&choice)
					if wantsToRoll(choice) {
						b[currentPlayer.startX][currentPlayer.startY] = "0"
						currentPlayer.posX, currentPlayer.posY, currentPlayer.laps = playerMove(currentPlayer)
						currentPlayer.totalLaps += currentPlayer.laps
						fmt.Println("lap: ", currentPlayer.laps)
						fmt.Println("total: ", currentPlayer.totalLaps)
						fmt.Println(currentPlayer.posX, currentPlayer.posY)
					}
				}
			} else {
				currentPlayer.posX, currentPlayer.posY, currentPlayer.laps = playerMove(currentPlayer)
				currentPlayer.totalLaps += currentPlayer.laps
				lapAction(currentPlayer.totalLaps, currentPlayer)
				gameEnded, winner = checkWin(currentPlayer.totalLaps, getPlayerName(currentPlayer))
				fmt.Println("lap: ", currentPlayer.laps)
				fmt.Println("total: ", currentPlayer.totalLaps)
				fmt.Println(currentPlayer.posX, currentPlayer.posY)
			}
			currentPlayer, otherPlayer = otherPlayer, currentPlayer
		} else {
			fmt.Println("Skipping turn.")
			currentPlayer, otherPlayer = otherPlayer, currentPlayer
		}
	}

	fmt.Println("Winner is: ", winner)
}
