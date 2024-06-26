package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/andrefsilveira1/graphql/graph"
	"github.com/andrefsilveira1/graphql/internal/database"
	_ "github.com/mattn/go-sqlite3"
)

const defaultPort = "8000"

func main() {
	fmt.Println("Started")
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	courseDB := database.NewCourse(db)
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		CategoryDB: categoryDB,
		CourseDB:   courseDB,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ ", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
