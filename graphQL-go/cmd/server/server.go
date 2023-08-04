package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/devfullcycle/13-GraphQL/graph"
	"github.com/devfullcycle/13-GraphQL/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8080"

func main() {
	// connect with database
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatalf("failed to open database:  %v", err)
	}
	defer db.Close()

	// create dbs
	categoryDb := database.NewCategory(db)
	courseDb := database.NewCourse(db)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		// inject categorydb into resolver
		CategoryDB: categoryDb,
		CourseDB:   courseDb,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
