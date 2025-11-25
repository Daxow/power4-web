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
