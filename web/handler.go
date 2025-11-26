package web

import (
	"log"
	"net/http"
	"power4/game"
)

type Server struct {
	Game *game.Game
}

func NewServer() *Server {
	g := game.NewGame()
	return &Server{
		Game: &g,
	}
}

func (s *Server) RegisterRoutes() {
	http.HandleFunc("/", s.handleIndex)
}

func (s *Server) handleIndex(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Puissance 4"))
	if err != nil {
		log.Println("write error:", err)
	}
}
