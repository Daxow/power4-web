package game

type Game struct {
	Board         [6][7]string
	CurrentPlayer int
	Winner        int
	GameOver      bool
	Message       string
}

func NewGame() Game {
	g := Game{}
	g.CurrentPlayer = 1
	g.Message = "Tour du Joueur 1"
	return g
}

func (g *Game) Reset() {
	*g = NewGame()
}

func (g *Game) DropPiece(col int) int {
	rowPlaced := -1
	for row := len(g.Board) - 1; row >= 0; row-- {
		if g.Board[row][col] == "" {
			symbol := "X"
			if g.CurrentPlayer == 2 {
				symbol = "O"
			}
			g.Board[row][col] = symbol
			rowPlaced = row
			break
		}
	}
	return rowPlaced
}

func (g *Game) SwitchPlayer() {
	if g.CurrentPlayer == 1 {
		g.CurrentPlayer = 2
	} else {
		g.CurrentPlayer = 1
	}
}
