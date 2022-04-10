package main

import (
	"log"
	"net/http"
	"os"

	"antonio.trupac/config"
	"antonio.trupac/directives"
	"antonio.trupac/graph"
	"antonio.trupac/graph/generated"
	"antonio.trupac/middlewares"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
)

const defaultPort = "8080"

func main() {
	config.MigrateTable()

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	// defer closing database
	db := config.GetDB()
	sqlDB, _ := db.DB()
	defer sqlDB.Close()

	router := mux.NewRouter()
	router.Use(middlewares.AuthMiddleware)

	c := generated.Config{Resolvers: &graph.Resolver{}}
	c.Directives.Auth = directives.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
