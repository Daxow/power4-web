package web

import (
    "html/template"
    "log"
    "net/http"
    "strconv"
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

    err := r.ParseForm()
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    colStr := r.FormValue("column")
    _, err = strconv.Atoi(colStr)
    if err != nil {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}
