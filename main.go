package main

import (
	"awesomeProject/transport/http/routes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"os/signal"
)

func main() {
	log.Fatal(fmt.Sprintf("Service shut down: %s", run()))
}

func run() error {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	_, cancel := context.WithCancel(context.Background())
	defer cancel()
	gracefulShutdown(cancel)

	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	log.Fatal(router.Run(":" + port))
	return nil
}

func gracefulShutdown(c context.CancelFunc) {
	osC := make(chan os.Signal, 1)
	signal.Notify(osC, os.Interrupt)
	go func() {
		log.Print(<-osC)
	}()
}
