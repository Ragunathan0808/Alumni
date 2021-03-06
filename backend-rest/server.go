package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func server(router *gin.Engine) {

	HOST := os.Getenv("HOST")

	server := http.Server{
		Addr:    HOST,
		Handler: router,
	}

	//Listen and serve go-routene
	go func() {
		log.Println("Stating Server on host:", HOST)
		err := server.ListenAndServe()
		if err != nil {
			log.Fatalln(err)
		}
	}()
	// Create and register close signal
	closeSignal := make(chan os.Signal)
	//signal.Notify(closeSignal)
	// Wait for close signal, then gracefull shutdown server
	<-closeSignal
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Server Shutdown...")

}
