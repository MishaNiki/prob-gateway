package server

import (
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// MainServer ...
type MainServer struct {
	server *http.Server
}

// NewMainServer
func NewMainServer() *MainServer {
	return &MainServer{}
}

// Start ...
func (ms *MainServer) Start() error {
	return ms.server.ListenAndServe()
}

// Configure ...
func (ms *MainServer) Configure() error {

	// ctx := context.Background()

	mainRouter := mux.NewRouter()
	// mainRouter.Handle("/echo/", echoRouter)
	// mainRouter.Handle("/hello/", helloRouter)

	// CORS polici
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Token"})
	method := handlers.AllowedMethods([]string{"POST", "GET", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})

	ms.server = &http.Server{
		Addr:    ":8081",
		Handler: handlers.CORS(header, method, origins)(mainRouter),
	}

	return nil
}

// https://medium.com/@Mark.io/rest-communication-to-containerd-recipe-plus-aws-api-gateway-sprinkled-on-top-25ef1011bbf0
// https://programmaticponderings.com/category/software-development/client-side-development/
