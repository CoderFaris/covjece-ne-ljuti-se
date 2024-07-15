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
	PlayerOut      bool
	PosX, PosY     int
	TotalLaps      int
	PathIndex      int
	Path           [][]int
	StartX, StartY int
	Laps           int
	Color          string
}

var Player1Path = [][]int{

	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0},
	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 10},
	{6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
}

var Player2Path = [][]int{

	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 10},
	{6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 0},
}

var Player3Path = [][]int{

	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 10},
	{6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 0},
	{4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
}

var Player4Path = [][]int{

	{6, 10}, {6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
	{8, 6}, {9, 6}, {10, 6}, {10, 5},
	{10, 4}, {9, 4}, {8, 4}, {7, 4}, {6, 4},
	{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 0},
	{4, 1}, {4, 2}, {4, 3}, {4, 4},
	{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5},
	{0, 6}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
	{4, 8}, {4, 9}, {4, 10}, {5, 10},
}

var Red = color.New(color.FgRed).SprintFunc()
var Blue = color.New(color.FgBlue).SprintFunc()
var Yellow = color.New(color.FgYellow).SprintFunc()
var Black = color.New(color.FgBlack).SprintFunc()

var player1 = Player{
	PosX:      10,
	PosY:      4,
	Path:      Player1Path,
	StartX:    10,
	StartY:    4,
	PlayerOut: false,
	Color:     Yellow("1"),
}

var player2 = Player{
	PosX:      4,
	PosY:      0,
	Path:      Player2Path,
	StartX:    4,
	StartY:    0,
	PlayerOut: false,
	Color:     Red("1"),
}

var player3 = Player{

	PosX:      0,
	PosY:      6,
	Path:      Player3Path,
	StartX:    0,
	StartY:    6,
	PlayerOut: false,
	Color:     Blue("1"),
}

var player4 = Player{

	PosX:      6,
	PosY:      10,
	Path:      Player4Path,
	StartX:    6,
	StartY:    10,
	PlayerOut: false,
	Color:     Black("1"),
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
	board[0][0] = Red("0")
	board[0][1] = Red("0")
	board[1][0] = Red("0")
	board[1][1] = Red("0")

	board[0][9] = Blue("0")
	board[0][10] = Blue("0")
	board[1][9] = Blue("0")
	board[1][10] = Blue("0")

	board[9][0] = Yellow("0")
	board[9][1] = Yellow("0")
	board[10][0] = Yellow("0")
	board[10][1] = Yellow("0")

	board[9][9] = Black("0")
	board[9][10] = Black("0")
	board[10][9] = Black("0")
	board[10][10] = Black("0")

	// home (ending) positions
	board[1][5] = Blue("h")
	board[2][5] = Blue("h")
	board[3][5] = Blue("h")
	board[4][5] = Blue("h")

	board[6][5] = Yellow("h")
	board[7][5] = Yellow("h")
	board[8][5] = Yellow("h")
	board[9][5] = Yellow("h")

	board[5][1] = Red("h")
	board[5][2] = Red("h")
	board[5][3] = Red("h")
	board[5][4] = Red("h")

	board[5][6] = Black("h")
	board[5][7] = Black("h")
	board[5][8] = Black("h")
	board[5][9] = Black("h")

	// starting positions
	board[10][4] = Yellow("A")
	board[4][0] = Red("A")
	board[0][6] = Blue("A")
	board[6][10] = Black("A")

	return board
}

// current board state
func DrawCurrentBoard(board *Board) {
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

	if totalLaps == 4 && (player == "P1" || player == "P2" || player == "P3" || player == "P4") {

		return true, player

	}

	return false, ""

}

func PlayerOut(player *Player, b *Board) bool {

	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)
	if !isOut(roll) {

		fmt.Println("Roll must be a six")
		return false

	}

	// player managed to go to the starting point
	player.PlayerOut = true
	(*b)[player.StartX][player.StartY] = player.Color
	DrawCurrentBoard(b)

	return true

}

func PlayerMove(player *Player, players []*Player, b *Board) (int, int, int) {
	fmt.Print("Rolling dice...\n")
	time.Sleep(2 * time.Second)
	roll := RollDice()
	fmt.Println(roll)

	newIndex, laps := getNextPosition(player.Path, player.PathIndex, roll)
	newPosX := player.Path[newIndex][0]
	newPosY := player.Path[newIndex][1]

	fmt.Printf("Player %v moving from (%d, %d) to (%d, %d)\n", getPlayerName(player), player.PosX, player.PosY, newPosX, newPosY)
	piece := b[newPosX][newPosY]
	fmt.Printf("Piece is: %v\n", piece)

	if EatPiece(newPosX, newPosY, b) {
		piece := (*b)[newPosX][newPosY]
		fmt.Printf("Detected piece at (%d, %d): %v\n", newPosX, newPosY, piece)
		otherPlayer := findPlayer(piece, players)

		if otherPlayer != nil {
			// resetting player who is eaten
			otherPlayer.PlayerOut = false
			otherPlayer.PosX = otherPlayer.StartX
			otherPlayer.PosY = otherPlayer.StartY
			otherPlayer.TotalLaps = 0
			otherPlayer.Laps = 0

			fmt.Printf("%v's piece was eaten and sent back to the out section.\n", getPlayerName(otherPlayer))
		} else {
			fmt.Println("Error: Could not find the player associated with the piece.")
		}
	}

	if isBooster(newPosX, newPosY, b) {
		fmt.Println("Entering booster chain")
		(*b)[player.PosX][player.PosY] = "0"

		// check chaining boosters
		for isBooster(newPosX, newPosY, b) {
			fmt.Printf("Booster activated at (%d, %d)\n", newPosX, newPosY)
			(*b)[newPosX][newPosY] = "0"

			boostIndex, boostLaps := getNextPosition(player.Path, newIndex, roll)
			laps += boostLaps
			newIndex = boostIndex

			if newIndex >= len(player.Path) {
				newIndex = newIndex % len(player.Path)
				laps++
			}

			newPosX = player.Path[newIndex][0]
			newPosY = player.Path[newIndex][1]

			fmt.Printf("Moved to (%d, %d)\n", newPosX, newPosY)

			if EatPiece(newPosX, newPosY, b) {
				piece := (*b)[newPosX][newPosY]
				fmt.Printf("Detected piece at (%d, %d): %v\n", newPosX, newPosY, piece)
				otherPlayer := findPlayer(piece, players)

				if otherPlayer != nil {
					otherPlayer.PlayerOut = false
					otherPlayer.PosX = otherPlayer.StartX
					otherPlayer.PosY = otherPlayer.StartY
					otherPlayer.TotalLaps = 0
					otherPlayer.Laps = 0

					fmt.Printf("%v's piece was eaten and sent back to the out section.\n", getPlayerName(otherPlayer))
				} else {
					fmt.Println("Error: Could not find the player associated with the piece.")
				}
			}
		}

		// place the player on the final non-booster tile
		(*b)[newPosX][newPosY] = player.Color
		player.PosX, player.PosY, player.PathIndex = newPosX, newPosY, newIndex
		fmt.Printf("Booster chain ended. Final position: (%d, %d)\n", newPosX, newPosY)
		DrawCurrentBoard(b)
		return newPosX, newPosY, laps
	}

	if isValidPosition(newPosX, newPosY, b) {
		fmt.Println("entering isvalid")
		(*b)[player.PosX][player.PosY] = "0"
		(*b)[newPosX][newPosY] = player.Color
		DrawCurrentBoard(b)
		player.PosX, player.PosY, player.PathIndex = newPosX, newPosY, newIndex
		return newPosX, newPosY, laps
	} else {
		fmt.Println("Invalid position")
		(*b)[player.PosX][player.PosY] = player.Color
		DrawCurrentBoard(b)
		return player.PosX, player.PosY, 0
	}

}

func getNextPosition(path [][]int, currentIndex int, roll int) (int, int) {

	newPosition := currentIndex + roll
	laps := newPosition / len(path)
	newPosition = newPosition % len(path)

	return newPosition, laps

}

func isValidPosition(posX, posY int, b *Board) bool {
	if posX < 0 || posX >= len(b) || posY < 0 || posY >= len(b[0]) {
		return false
	}
	cell := b[posX][posY]
	return cell != "#" && cell != "h"
}

func EatPiece(posX, posY int, b *Board) bool {

	fmt.Printf("Checking piece at (%d, %d)\n", posX, posY)
	if b[posX][posY] == Yellow("1") || b[posX][posY] == Red("1") || b[posX][posY] == Blue("1") || b[posX][posY] == Black("1") {

		return true

	}

	return false

}

func isBooster(posX, posY int, b *Board) bool {

	return b[posX][posY] == "B"

}

func generateBoosters(b *Board) {

	// 4 boosters

	positions := [][]int{{9, 4}, {8, 4}, {7, 4}, {6, 4},
		{6, 3}, {6, 2}, {6, 1}, {6, 0}, {5, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
		{3, 4}, {2, 4}, {1, 4}, {0, 4}, {0, 5}, {1, 6}, {2, 6}, {3, 6}, {4, 6}, {4, 7},
		{4, 8}, {4, 9}, {4, 10}, {5, 10}, {6, 9}, {6, 8}, {6, 7}, {6, 6}, {7, 6},
		{8, 6}, {9, 6}, {10, 6}, {10, 5}}

	var randomPos [][]int
	var i int

	for i < 4 {

		randomPos = append(randomPos, positions[rand.IntN(len(positions))])
		i += 1
	}

	b[randomPos[0][0]][randomPos[0][1]] = "B"
	b[randomPos[1][0]][randomPos[1][1]] = "B"
	b[randomPos[2][0]][randomPos[2][1]] = "B"
	b[randomPos[3][0]][randomPos[3][1]] = "B"

}

func findPlayer(piece any, players []*Player) *Player {
	fmt.Println("Finding player for piece: ", piece)
	for _, player := range players {
		fmt.Printf("Checking player: %v with color: %v\n", getPlayerName(player), player.Color)
		if player.Color == piece {
			fmt.Printf("Found matching player: %v\n", getPlayerName(player))
			return player
		}
	}
	fmt.Println("No player found for piece: ", piece)
	return nil
}

func lapAction(lap int, player *Player, b *Board) {

	switch lap {

	case 1:

		if player == &player1 {

			b[9][5] = "1"

		}

		if player == &player2 {

			b[5][1] = "1"

		}

		if player == &player3 {

			b[1][5] = "1"

		}

		if player == &player4 {

			b[5][9] = "1"

		}

	case 2:

		if player == &player1 {

			b[8][5] = "1"

		}

		if player == &player2 {

			b[5][2] = "1"

		}

		if player == &player3 {

			b[2][5] = "1"

		}

		if player == &player4 {

			b[5][8] = "1"

		}

	case 3:

		if player == &player1 {

			b[7][5] = "1"

		}

		if player == &player2 {

			b[5][3] = "1"

		}

		if player == &player3 {

			b[3][5] = "1"

		}

		if player == &player4 {

			b[5][7] = "1"

		}

	case 4:

		if player == &player1 {

			b[6][5] = "1"

		}

		if player == &player2 {

			b[5][4] = "1"

		}

		if player == &player3 {

			b[4][5] = "1"

		}

		if player == &player4 {

			b[5][6] = "1"

		}

	}

}

func getPlayerName(player *Player) string {

	if player == &player1 {

		return "P1"

	}

	if player == &player2 {

		return "P2"
	}

	if player == &player3 {

		return "P3"
	}

	return "P4"

}

func resetTurn(turn int) int {

	if turn > 3 {

		turn = 0

	}

	return turn

}

func Game() {

	var players = []*Player{&player1, &player2, &player3, &player4}
	var turn int = 0

	var choice int
	var gameEnded = false
	var winner string

	b := GetBoard()

	b[9][4] = "B"

	generateBoosters(&b)
	DrawCurrentBoard(&b)

	for !gameEnded {
		fmt.Printf("(%v) Roll dice? (1 for Yes, any key for No): ", getPlayerName(players[turn]))
		fmt.Scan(&choice)
		if wantsToRoll(choice) {
			if !players[turn].PlayerOut {
				if PlayerOut(players[turn], &b) {
					fmt.Printf("(%v) Roll again? (1 for Yes, any key for No): ", getPlayerName(players[turn]))
					fmt.Scan(&choice)
					if wantsToRoll(choice) {
						b[players[turn].StartX][players[turn].StartY] = "0"
						players[turn].PosX, players[turn].PosY, players[turn].Laps = PlayerMove(players[turn], players, &b)
						players[turn].TotalLaps += players[turn].Laps
						fmt.Println("lap: ", players[turn].Laps)
						fmt.Println("total: ", players[turn].TotalLaps)
						fmt.Println(players[turn].PosX, players[turn].PosY)
					}
				}
			} else {
				players[turn].PosX, players[turn].PosY, players[turn].Laps = PlayerMove(players[turn], players, &b)
				players[turn].TotalLaps += players[turn].Laps
				lapAction(players[turn].TotalLaps, players[turn], &b)
				gameEnded, winner = checkWin(players[turn].TotalLaps, getPlayerName(players[turn]))
				fmt.Println("lap: ", players[turn].Laps)
				fmt.Println("total: ", players[turn].TotalLaps)
				fmt.Println(players[turn].PosX, players[turn].PosY)
			}
			turn += 1
			turn = resetTurn(turn)

		} else {
			fmt.Println("Skipping turn.")
			turn += 1
			turn = resetTurn(turn)
		}
	}

	fmt.Println("Winner is: ", winner)

}
