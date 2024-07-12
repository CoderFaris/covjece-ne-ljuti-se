package game_test

import (
	"fmt"
	"testing"

	"github.com/CoderFaris/dontgetmad/internal/game"
)

func TestEatPiece(t *testing.T) {

	player1 := &game.Player{
		PlayerOut: true,
		PosX:      4,
		PosY:      3,
		Color:     game.Yellow("1"),
		Path:      game.Player1Path,
		StartX:    10,
		StartY:    4,
		PathIndex: 13,
	}

	player2 := &game.Player{
		PlayerOut: true,
		PosX:      4,
		PosY:      4,
		Color:     game.Red("1"),
		Path:      game.Player2Path,
		StartX:    4,
		StartY:    0,
		PathIndex: 4,
	}

	players := []*game.Player{player1, player2}

	b := game.GetBoard()

	b[player1.PosX][player1.PosY] = player1.Color
	b[player2.PosX][player2.PosY] = player2.Color

	// print board state before move
	game.DrawCurrentBoard(&b)
	fmt.Printf("Initial board positions:\n")
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Printf("%s ", b[i][j])
		}
		fmt.Println()
	}

	fmt.Printf("Player1 (%s) at (%d, %d)\n", player1.Color, player1.PosX, player1.PosY)
	fmt.Printf("Player2 (%s) at (%d, %d)\n", player2.Color, player2.PosX, player2.PosY)

	player1.PosX, player1.PosY, _ = game.PlayerMove(player1, players, &b)

	// print board state after move
	fmt.Printf("Board positions after move:\n")
	for i := 0; i < len(b); i++ {
		for j := 0; j < len(b[0]); j++ {
			fmt.Printf("%s ", b[i][j])
		}
		fmt.Println()
	}

	// print player positions after move
	fmt.Printf("After move player1 at (%d, %d)\n", player1.PosX, player1.PosY)
	fmt.Printf("After move player2 at (%d, %d)\n", player2.PosX, player2.PosY)

	if player2.PosX != player2.StartX || player2.PosY != player2.StartY {
		t.Errorf("Expected player2 to be at start position, got (%d, %d)", player2.PosX, player2.PosY)
	}
}
