package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/gorilla/websocket"
	"github.com/minskylab/calab/experiments"
	"github.com/minskylab/calab/server/graph"
	"github.com/minskylab/calab/server/graph/generated"
	"github.com/rs/cors"
)

func ServeExperiment(exp *experiments.Experiment, port int) {
	router := chi.NewRouter()

	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8080", "http://localhost:3000"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	srv := handler.New(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		Exp: exp,
	}}))

	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			// EnableCompression: true,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})

	srv.Use(extension.Introspection{})

	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))
	router.Handle("/graphql", srv)

	portString := fmt.Sprintf(":%d", port)

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(portString, nil))

	err := http.ListenAndServe(portString, router)
	if err != nil {
		panic(err)
	}
}
