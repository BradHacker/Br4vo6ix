package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/BradHacker/Br4vo6ix/ent"
	"github.com/BradHacker/Br4vo6ix/graph"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	"github.com/sirupsen/logrus"
)

// Defining the Graphql handler
func graphqlHandler(client *ent.Client) gin.HandlerFunc {
	// NewExecutableSchema and Config are in the generated.go file
	// Resolver is in the resolver.go file
	h := handler.NewDefaultServer(graph.NewSchema(client))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Defining the Playground handler
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func RunAPI(client *ent.Client, wg *sync.WaitGroup) {
	defer wg.Done()

	client = ent.SQLLiteOpen("file:br4vo6ix.sqlite?_loc=auto&cache=shared&_fk=1")
	fmt.Println("Local SQLite Database Initialized...")

	defer client.Close()
	ctx := context.Background()

	// Auto migrate the database
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	corsAllowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if corsAllowedOrigins == "" {
		corsAllowedOrigins = "*"
	}
	allowedOrigins := strings.Split(corsAllowedOrigins, ",")

	r := gin.Default()

	// Cors magic 🤩
	r.Use(cors.New(cors.Config{
		AllowOrigins:     allowedOrigins,
		AllowMethods:     []string{"GET", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	r.POST("/query", graphqlHandler(client))
	r.GET("/playground", playgroundHandler())
	if err := r.Run(port); err != nil {
		logrus.Errorf("failed to start API server: %v", err)
	}
}
