package main

import (
	"net/http"
	"time"
	"video_server/router"
)

func main() {
	server := &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      router.InitRouter(),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}
