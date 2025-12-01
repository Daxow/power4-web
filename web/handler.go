package web

import (
    "html/template"
    "log"
    "net/http"

    "power4/game"
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
