package main

import (
	"Bot-or-Not/internal/di"
	"Bot-or-Not/pkg/config"
	"log"
	"net/http"

	_ "Bot-or-Not/internal/migration"
)

func main() {
	root := di.New()

	if err := http.ListenAndServe(":"+config.PORT, root.Echo); err != nil {
		log.Fatal("Error starting server: ", err)
	}

}
