package main

import (
	"fmt"
	"net/http"

	"github.com/jroyal/atx-metro/api"
)

func main() {
	fmt.Print("Loading feed info...")
	service := api.NewService(&api.MetroServiceConfig{
		FeedDirectory: "/Users/james/Downloads/capmetro",
	})
	fmt.Println(" done.")
	http.ListenAndServe("localhost:3001", service.Handler)
}
