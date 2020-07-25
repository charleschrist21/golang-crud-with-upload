package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : this is helper function connect mongoDB
// if you want to export your function. You must to start upercase function name.
// Otherwise you won't see your function when you import other class
func ConnectDB() *mongo.Collection {
	//Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost")

	// connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connect to MongoDB")

	collection := client.Database("go_rest_api").Collection("books")

	return collection

}

// ErrorResponse : this is error model
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : this is helper function to prepare error model
func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
