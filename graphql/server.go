package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/shivamk2406/Practice/configs"
	"github.com/shivamk2406/Practice/database"
	"github.com/shivamk2406/Practice/graphql/graph"
	"github.com/shivamk2406/Practice/graphql/graph/generated"
	"github.com/shivamk2406/Practice/internal/service/user"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	conf := configs.LoadAppConfig()

	db, clean, err := database.Open(conf)
	if err != nil {
		log.Println(err)
	}

	userDB := user.NewRepo(db)
	resolver := graph.NewResolver(userDB)
	defer clean()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
