package main

import (
	"log"

	"github.com/deerbone/killgrave/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
