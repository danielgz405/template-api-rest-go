package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BacataCode/SmartCardConnectApiRest/handlers"
	"github.com/BacataCode/SmartCardConnectApiRest/server"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	JWT_SECRET := os.Getenv("JWT_SECRET")
	MONGO_TEST := os.Getenv("DB_URI_TEST")
	TESTING_MODE := os.Getenv("TESTING_MODE") == "true"

	DB_URI := os.Getenv("DB_URI")
	if TESTING_MODE {
		DB_URI = MONGO_TEST
		fmt.Println("Ⓐ ☭------------♥♥♥ THE MODE IS TESTING ♥♥♥------------☭ Ⓐ")
	}

	s, err := server.NewServer(context.Background(), &server.Config{
		Port:      ":" + PORT,
		JWTSecret: JWT_SECRET,
		DbURI:     DB_URI,
	})
	if err != nil {
		log.Fatal(err)
	}

	s.Start(BindRoutes)

}

func BindRoutes(s server.Server, r *mux.Router) {
	r.HandleFunc("/welcome/{name}", handlers.HomeHandler(s)).Methods(http.MethodGet)

	//Auth
	r.HandleFunc("/login", handlers.LoginHandler(s)).Methods(http.MethodPost)

	//user
	r.HandleFunc("/user/create", handlers.CreateUserHandler(s)).Methods(http.MethodPost)
	r.HandleFunc("/user/delete/{id}", handlers.DeleteUserHandler(s)).Methods(http.MethodDelete)
	r.HandleFunc("/user/update/{id}", handlers.UpdateAnyUserHandler(s)).Methods(http.MethodPatch)
	r.HandleFunc("/users/list", handlers.ListUsersHandler(s)).Methods(http.MethodGet)
	r.HandleFunc("/user/profile", handlers.ProfileHandler(s)).Methods(http.MethodGet)

	//WS
	r.HandleFunc("/ws/{Authorization}/{Module}", s.Hub().HandleWebSocket(s.Config().JWTSecret))
}
