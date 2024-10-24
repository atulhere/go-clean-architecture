package main

import (
	"database/sql"
	"go-clean-architecture/handler"
	mysql "go-clean-architecture/infrastructure"
	"go-clean-architecture/usecase"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	dsn := "root:@tcp(127.0.0.1:3306)/inventory"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// Initialize repository, usecase, and handler
	userRepo := mysql.NewMySQLUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// HTTP routes
	http.HandleFunc("POST /login", userHandler.LoginHandler)
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
