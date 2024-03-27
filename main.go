package main

import (
	"fmt"
	"log"
)

const NoWinner = "none"

var switchPlayer = false

var colors = map[string][2]string{
	"orange": {"\x1b[34m", "\x1b[0m"},
	"cyan":   {"\x1b[36m", "\x1b[0m"},
	"black":  {"\x1b[30m", "\x1b[0m"},
}

func colorize(text string, color string) string {
	return fmt.Sprintf("%s%s%s", colors[color][0], text, colors[color][1])
}

type boardType map[int]string

type Player struct {
	Symbol string
	Score  int
}

var playerX = Player{colorize("x", "orange"), 0}
var playerO = Player{colorize("o", "cyan"), 0}

var players = []Player{
	playerX,
	playerO,
}

func getPlayer(ps []Player) Player {
	switchPlayer = !switchPlayer
	if switchPlayer {
		return ps[0]
	} else {
		return ps[1]
	}
}

func printBoard(b boardType) {
	for i := 0; i < 9; i += 3 {
		fmt.Printf("%v | %v | %v\n", colorize(b[i], "black"), colorize(b[i+1], "black"), colorize(b[i+2], "black"))
	}
}

func clearOutput() {
	for i := 0; i <= 4; i++ {
		fmt.Printf("\033[1A\033[K")
	}
}

func isPosOutOfRange(pos int) bool {
	if pos < 0 || pos > 8 {
		return true
	}
	return false
}

func isPosChecked(pos int, b boardType) bool {
	if b[pos] == colorize("x", "orange") || b[pos] == colorize("o", "cyan") {
		return true
	}
	return false
}

func play(b boardType, ps []Player) boardType {
	var pos int
	player := getPlayer(ps)

	fmt.Printf("Score: (%v:%v, %v:%v) - turn: %v\n", ps[0].Symbol, ps[0].Score, ps[1].Symbol, ps[1].Score, player.Symbol)
	_, err := fmt.Scanln(&pos)

	if err != nil {
		log.Fatal("Not a valid position ", err)
	}

	if isPosChecked(pos, b) || isPosOutOfRange(pos) {
		return b
	}

	b[pos] = player.Symbol
	return b
}

func getWinner(b boardType, ps []Player) string {
	winCombos := [][]int{
		{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // verticals
		{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // horizontals
		{0, 4, 8}, {2, 4, 6}, // diagonals
	}

	for _, combo := range winCombos {
		if b[combo[0]] == b[combo[1]] && b[combo[1]] == b[combo[2]] {
			if b[combo[0]] == ps[0].Symbol {
				ps[0].Score += 1
			} else {
				ps[1].Score += 1
			}
			return b[combo[0]]
		}
	}

	return NoWinner
}

func initBoard() boardType {
	board := make(boardType)
	for i := 0; i < 9; i++ {
		board[i] = colorize(fmt.Sprintf("%d", i), "black")
	}
	return board
}

func (b boardType) isFull() bool {
	for _, v := range b {
		if v != colorize("x", "orange") && v != colorize("o", "cyan") {
			return false
		}
	}
	return true
}

func shouldReset(w string, b boardType) bool {
	if w != NoWinner || (w == NoWinner && b.isFull()) {
		return true
	}
	return false
}

func main() {
	b := initBoard()
	for {
		printBoard(b)
		b = play(b, players)
		winner := getWinner(b, players)
		clearOutput()
		if shouldReset(winner, b) {
			b = initBoard()
		}
	}
}
