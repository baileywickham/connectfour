package main

import "bufio"
import "errors"
import "os"
import "strconv"
import "strings"

const totalRows = 6
const totalColumns = 7

type Game struct {
	Board      [totalRows][totalColumns]string
	Turn       string
	TurnNumber int
}

func (g *Game) addTile(row int) error {

	for i := 0; i < totalRows; i++ {
		if g.Board[i][row] == "-" {
			g.Board[i][row] = g.Turn
			g.checkForWin(row, i)
			g.nextTurn()
			return nil
		}
	}
	return errors.New("Full")
}

func (g *Game) checkForWin(row, column int) {
	if g.countHorizontal(g.Turn, column) || g.countVertical(g.Turn, row) {
		panic("win")
	}
}

func (g *Game) nextTurn() {
	g.TurnNumber++
	if g.Turn == "X" {
		g.Turn = "O"
	} else {
		g.Turn = "X"
	}

}

func (g *Game) countHorizontal(token string, row int) bool {
	counter := 0
	for i := 0; i < totalColumns; i++ {
		if g.Board[row][i] == token {
			counter++
			if counter == 4 {
				return true

			}
		} else {
			counter = 0
		}
	}
	return false
}

func (g *Game) countVertical(token string, column int) bool {
	counter := 0
	for i := 0; i < totalRows; i++ {
		if g.Board[i][column] == token {
			counter++
			if counter == 4 {
				return true

			}
		} else {
			counter = 0
		}
	}
	return false
}

func newBoard() *Game {
	g := Game{Turn: "X"}
	for i := 0; i < totalRows; i++ {
		for j := 0; j < totalColumns; j++ {
			g.Board[i][j] = "-"
		}
	}

	return &g
}

func (g *Game) printBoard() {
	for i := totalRows - 1; i >= 0; i-- {
		for j := 0; j < totalColumns; j++ {
			print(g.Board[i][j], " ")
		}
		println()
	}
	println()
}

func input() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, err := reader.ReadString('\n')
		char = strings.TrimSuffix(char, "\n")
		row, err := strconv.Atoi(char)
		if err != nil {
			panic(err)
		}

		if row >= 0 && row < totalRows {
			return row
		} else {
			println("invalid input", row)
		}
	}
}

func main() {
	runner()
}

func runner() {
	g := newBoard()
	for {
		println(g.Turn, "'s turn")
		row := input()
		if err := g.addTile(row); err != nil {
			println("Invalid Input")
		}
		g.printBoard()
	}
}
