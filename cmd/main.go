package main

import (
	"github.com/RianIhsan/go-elastic-query-common/internal/http/delivery"
	"github.com/RianIhsan/go-elastic-query-common/internal/repository/elasticsearch"
	"github.com/RianIhsan/go-elastic-query-common/internal/usecase"
	"github.com/olivere/elastic/v7"
	"log"
	"net/http"
)

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9700"))
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}

	userRepo := elasticsearch.NewUserRepository(client)
	userUC := usecase.NewUseCase(userRepo)
	userHandler := delivery.NewUserHandler(userUC)

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/user", userHandler.GetUser)

	log.Println("Server is running at :8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
}
