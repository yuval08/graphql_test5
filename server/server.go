package main

import (
	"fmt"
	"github.com/99designs/gqlgen/handler"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"log"
	"net/http"
	"training/graphql_test5"
	"training/graphql_test5/postgres"
)

const (
	port       = 8080
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "docker"
	dbName     = "testgql"
)

func main() {
	router, db := initialize()
	defer db.Close()

	log.Printf("connect to http://localhost:%d/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}

func initialize() (*chi.Mux, *postgres.Db) {
	router := chi.NewRouter()

	db, err := postgres.New(postgres.ConnString(dbHost, dbPort, dbUser, dbPassword, dbName))

	if err != nil {
		log.Fatal(err)
	}

	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,          //log api request calls
		middleware.DefaultCompress, //compress results, mostly gzipping assets and json
		middleware.StripSlashes,    //match paths with a training slash, string it and continue router through the mux
		middleware.Recoverer,       //recover from panic without crashing server
	)

	router.Handle("/", handler.Playground("GraphQL playground", "/graphql"))
	router.Post("/graphql", handler.GraphQL(graphql_test5.NewExecutableSchema(graphql_test5.Config{Resolvers: &graphql_test5.Resolver{Database: db}})))

	return router, db
}
