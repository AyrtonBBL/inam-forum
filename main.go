package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Bienvenue sur I Need a Mate </h1><p>Prêt à trouver ton futur mate ?</p>")
	})

	fmt.Println("Serveur lancé sur http://localhost:8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
