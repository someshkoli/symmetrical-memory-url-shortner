package main

import (
	"fmt"

	"github.com/someshkoli/symmetrical-memory-url-shorner/pkg/api"
	"github.com/someshkoli/symmetrical-memory-url-shorner/pkg/store"
)

func main() {
	urlshortner := api.NewURLShortner(store.NewFileRecordStorage("store", 10), 8888, "localhost")
	fmt.Println("Starting Url Shortner server")
	if urlshortner.Start() != nil {
		fmt.Println("Error starting server")
	}
}
