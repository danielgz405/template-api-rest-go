package server

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/BacataCode/SmartCardConnectApiRest/database"
	"github.com/BacataCode/SmartCardConnectApiRest/repository"
	"github.com/BacataCode/SmartCardConnectApiRest/websocket"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type Config struct {
	Port      string
	JWTSecret string
	DbURI     string
}

type Server interface {
	Config() *Config
	Hub() *websocket.Hub
}

type Broker struct {
	config *Config
	router *mux.Router
	hub    *websocket.Hub
}

func (b *Broker) Config() *Config {
	return b.config
}

func (b *Broker) Hub() *websocket.Hub {
	return b.hub
}

func NewServer(ctx context.Context, config *Config) (*Broker, error) {
	if config.Port == "" {
		return nil, errors.New("port is required")
	}
	if config.JWTSecret == "" {
		return nil, errors.New("jwt secret is required")
	}
	if config.DbURI == "" {
		return nil, errors.New("database uri is required")
	}
	broker := &Broker{
		config: config,
		router: mux.NewRouter(),
		hub:    websocket.NewHub(),
	}
	return broker, nil
}

func (b *Broker) Start(binder func(s Server, r *mux.Router)) {
	b.router = mux.NewRouter()
	binder(b, b.router)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedHeaders:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PATCH", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	})

	handler := c.Handler(b.router)
	repo, err := database.NewMongoRepo(b.config.DbURI)
	if err != nil {
		log.Fatal(err)
	}

	go b.Hub().Run()

	repository.SetRepository(repo)

	log.Println("Server started on port", b.config.Port)
	if err := http.ListenAndServe(b.config.Port, handler); err != nil {
		log.Fatal("Server failed to start", err)
	}
}
