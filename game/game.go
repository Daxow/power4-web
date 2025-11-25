package game

type Game struct {
	Board         [6][7]string
	CurrentPlayer int
	Winner        int
	GameOver      bool
	Message       string
}
