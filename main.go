package main

import (
	"fmt"
	"log"
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

var player1 = Player{colorize("[Player x]", "blue"), colorize("x", "blue")}
var player2 = Player{colorize("[Player o]", "cyan"), colorize("o", "cyan")}

func getPlayer(t int) Player {
	if t%2 == 0 {
		return player1
	} else {
		return player2
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

func playTurn(b map[int]string, t int) (map[int]string, int, error) {
	var pos int
	player := getPlayer(t)

	fmt.Printf("%v chose a position from 0 to 8:\n", player.Name)
	_, err := fmt.Scanln(&pos)

	if err != nil {
		return b, t, err
	}

	if isPosChecked(pos, b) || isPosOutOfRange(pos) {
		// just return whatever
		return b, t, nil
	}

	b[pos] = player.Value
	t += 1
	return b, t, nil
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
		var err error
		board, turn, err = playTurn(board, turn)
		if err != nil {
			log.Fatal("Not a valid position ", err)
		}
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
