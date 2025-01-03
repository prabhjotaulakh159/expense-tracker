package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabhjotaulakh159/expense-tracker/db"
	"github.com/prabhjotaulakh159/expense-tracker/router"
	"github.com/prabhjotaulakh159/expense-tracker/server"
)

// https://gin-gonic.com/docs/examples/graceful-restart-or-stop/
func main() {
	_router := router.GetRouterInstance()
	_server := server.GetServerInstance("localhost", 8080, _router)

	_db, err := db.GetGormInstance()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// need sql object to close connection
	sqlDb, err := _db.DB()
	if err != nil {
		log.Fatalf("Error getting SQL object: %v", err)
	}
	defer sqlDb.Close()

	go func() {
		if err := _server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	log.Println("Database closing is deferred...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := _server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
