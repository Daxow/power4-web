package web

import (
	"html/template"
	"log"
	"net/http"
	"power4/game"
	"strconv"
)

type Server struct {
	Game *game.Game
	Tmpl *template.Template
}

func NewServer() *Server {
	g := game.NewGame()
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}
	return &Server{
		Game: &g,
		Tmpl: t,
	}
}

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/", s.handleIndex)
	http.HandleFunc("/play", s.handlePlay)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	err := s.Tmpl.Execute(w, s.Game)
	if err != nil {
		http.Error(w, "Error rendering template", http.StatusInternalServerError)
		return
	}
}

func (s *Server) handlePlay(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if s.Game.GameOver {
		s.Game.Reset()
	}

	err := r.ParseForm()
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	colStr := r.FormValue("column")
	col, err := strconv.Atoi(colStr)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if col < 0 || col > 6 {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	rowPlaced := s.Game.DropPiece(col)
	if rowPlaced == -1 {
		s.Game.Message = "Cette colonne est pleine"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if s.Game.CheckWin(rowPlaced, col) {
		s.Game.GameOver = true
		s.Game.Winner = s.Game.CurrentPlayer
		s.Game.Message = "Le Joueur " + strconv.Itoa(s.Game.CurrentPlayer) + " a gagnÃ© !"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if s.Game.IsDraw() {
		s.Game.GameOver = true
		s.Game.Winner = 0
		s.Game.Message = "Match nul"
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	s.Game.SwitchPlayer()
	s.Game.Message = "Tour du Joueur " + strconv.Itoa(s.Game.CurrentPlayer)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
