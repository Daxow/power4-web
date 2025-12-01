package main

import (
	"log"
	"net/http"
	"power4/web"
)

func main() {
	server := web.NewServer()
	server.RegisterRoutes()

	log.Println("Serveur lancé : http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
