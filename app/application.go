package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	url = "127.0.0.1:8080"
)

var (
	router = mux.NewRouter()
)

func StartApplication() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    url,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

	log.Println(fmt.Sprintf("server started at url: %s", url))
}
