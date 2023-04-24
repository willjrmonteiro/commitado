package main

import (
	"log"
	"net/http"
	"os"

	"commitado/graphql"
	"commitado/graphql/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/rs/cors"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	allowedCorsOrigins := os.Getenv("ALLOWED_CORS_ORIGINS")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graphql.Resolver{}}))

	corsOptions := cors.Options{
		AllowedOrigins: []string{allowedCorsOrigins},
	}
	corsMiddleware := cors.New(corsOptions).Handler

	http.Handle("/", corsMiddleware(playground.Handler("GraphQL playground", "/query")))
	http.Handle("/query", corsMiddleware(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
