package main

import (
	"Bot-or-Not/internal/di"
	"log"
	"net/http"

	_ "Bot-or-Not/internal/migration"
)

func main() {
	root := di.New()

	if err := http.ListenAndServe(":8080", root.Echo); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
