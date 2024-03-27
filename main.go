package main

import (
	"fmt"
	"log"
)

const NoWinner = "none"
const BoardSize = 9
const OutputLines = 4

var colors = map[string][2]string{
	"blue":  {"\x1b[34m", "\x1b[0m"},
	"cyan":  {"\x1b[36m", "\x1b[0m"},
	"black": {"\x1b[30m", "\x1b[0m"},
}

func colorize(text string, color string) string {
	return fmt.Sprintf("%s%s%s", colors[color][0], text, colors[color][1])
}

type Player struct {
	Name  string
	Value string
}

var playerX = Player{colorize("[Player x]", "blue"), colorize("x", "blue")}
var playerO = Player{colorize("[Player o]", "cyan"), colorize("o", "cyan")}

func getPlayer(t int) Player {
	if t%2 == 0 {
		return playerX
	} else {
		return playerO
	}
}

func printBoard(b map[int]string) {
	for i := 0; i < BoardSize; i += 3 {
		fmt.Printf("%v | %v | %v\n", colorize(b[i], "black"), colorize(b[i+1], "black"), colorize(b[i+2], "black"))
	}
}

func clearOutput() {
	for i := 0; i <= OutputLines; i++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func isPosOutOfRange(pos int) bool {
	if pos < 0 || pos > 8 {
		return true
	}
	return false
}

func isPosChecked(pos int, b map[int]string) bool {
	if b[pos] == colorize("x", "blue") || b[pos] == colorize("o", "cyan") {
		return true
	}
	return false
}

func playTurn(b map[int]string, t int) (map[int]string, int) {
	var pos int
	player := getPlayer(t)

	fmt.Printf("%v chose a position from 0 to 8:\n", player.Name)
	_, err := fmt.Scanln(&pos)

	if err != nil {
		log.Fatal("Not a valid position ", err)
	}

	if isPosChecked(pos, b) || isPosOutOfRange(pos) {
		// just return whatever
		return b, t
	}

	b[pos] = player.Value
	return b, t + 1
}

func checkWinner(b map[int]string) string {
	// Winning combinations
	winCombos := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // columns
		{0, 4, 8}, {2, 4, 6}, // diagonals
	}

	for _, combo := range winCombos {
		if b[combo[0]] == b[combo[1]] && b[combo[1]] == b[combo[2]] {
			return b[combo[0]]
		}
	}

	return NoWinner
}

func initBoard() map[int]string {
	board := make(map[int]string)
	for i := 0; i < BoardSize; i++ {
		board[i] = colorize(fmt.Sprintf("%d", i), "black")
	}
	return board
}

func main() {
	board := initBoard()
	turn := 0
	for turn < BoardSize {
		printBoard(board)
		board, turn = playTurn(board, turn)
		winner := checkWinner(board)
		clearOutput()
		if winner != NoWinner {
			fmt.Printf("%v wins\n", winner)
			break
		}
	}

}
