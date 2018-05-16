package main

import (
	"log"

	"github.com/hectorrios/business_card/actions"
)

func main() {
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}
