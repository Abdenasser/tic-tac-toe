package main

import (
	"fmt"
)

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

func createPlayer(t int) Player {
	if t%2 == 0 {
		return Player{colorize("[Player x]", "blue"), colorize("x", "blue")}
	} else {
		return Player{colorize("[Player o]", "cyan"), colorize("o", "cyan")}
	}
}

func printBoard(b map[int]string) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%v | %v | %v\n", colorize(b[i], "black"), colorize(b[i+1], "black"), colorize(b[i+2], "black"))
	}
}

func clearOutput() {
	// assuming we print 4 lines at each iteration
	for i := 0; i < 5; i++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func playTurn(b map[int]string, t int) (map[int]string, int) {
	var input int
	player := createPlayer(t)

	fmt.Printf("%v chose a position from 0 to 8:\n", player.Name)
	_, err := fmt.Scanln(&input)

	if err != nil || b[input] == colorize("x", "blue") || b[input] == colorize("o", "cyan") {
		// just return whatever
		return b, t
	}

	b[input] = player.Value
	t += 1
	return b, t
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

	return "none"
}

func main() {
	board := make(map[int]string)
	for i := 0; i < 9; i++ {
		board[i] = colorize(fmt.Sprintf("%d", i), "black")
	}
	turn := 0
	for turn < 9 {
		printBoard(board)
		board, turn = playTurn(board, turn)
		winner := checkWinner(board)
		if winner != "none" {
			clearOutput()
			printBoard(board)
			fmt.Printf("%v wins\n", winner)
			return
		}
		clearOutput()
	}

}
