package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"samueldasilvadev.com/goreddit/postgres"
	"samueldasilvadev.com/goreddit/web"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	store, err := postgres.NewStore(os.Getenv("CONN_STRING"))
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
