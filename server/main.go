package main

import (
	"fmt"
	"net/http"
	"log"
	"context"
	"os"
	"os/signal"
	"syscall"
	"github.com/prabhjotaulakh159/expense-tracker/db"
	"github.com/prabhjotaulakh159/expense-tracker/controllers"
	"github.com/prabhjotaulakh159/expense-tracker/services"
	"github.com/prabhjotaulakh159/expense-tracker/repositories"

)

func main() {
	db, err := db.NewGormDB()
	if err != nil {
		log.Fatalf("connecting to database: %v", err)
	}
	defer func() {
		sqlDb, err := db.DB()
		if err != nil {
			log.Fatalf("closing database connection: %v", err)
		}
		if err := sqlDb.Close(); err != nil {
			log.Fatalf("closing database connection: %v", err)
		}
		log.Println("database connection closed")
	}()

	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	mux := http.NewServeMux()	
	mux.HandleFunc("POST /api/user/register", userController.Register)
	server := newServer(mux, "localhost", 3000)
	go startServer(server)
	log.Println(fmt.Sprintf("server is listening on %s", server.Addr))

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	stopServer(server)
	log.Println("server stopped successfully")
}

func newServer(mux *http.ServeMux, host string, port int) *http.Server {
	return &http.Server {
		Addr: fmt.Sprintf("%s:%d", host, port),
		Handler: mux,
	}
}

func startServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("starting server: %v", err)
	}
}

func stopServer(server *http.Server) {
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("stopping server: %v", err)
	}
}