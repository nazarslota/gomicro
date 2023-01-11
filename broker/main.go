package main

import (
	"fmt"
	"net/http"

	"github.com/nazarslota/gomicro/broker/handler"
)

const port = "8080"

func main() {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: handler.New().H(),
	}

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
