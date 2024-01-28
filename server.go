package main

import (
	"os"

	"github.com/amro-alasri/graphQL-server/controllers"
	"github.com/gin-gonic/gin"
)

const defaultPort = ":8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	server := gin.Default()

	server.GET("/", controllers.Playground())
	server.POST("/query", controllers.GraphQLHandler())
	server.Run(defaultPort)

	// srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	// http.Handle("/query", srv)

	// log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	// log.Fatal(http.ListenAndServe(":"+port, nil))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("what is your name"))
	// })
}
