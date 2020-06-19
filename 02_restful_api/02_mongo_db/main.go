// all credits to https://www.youtube.com/watch?v=oW7PMHEYiSk
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"` // bson is offical annotation for mongodb
	Firstname string             `json:"firstname,omitempty" bson:"firstname,omitempty"`
	Lastname  string             `json:"lastname,omitempty" bson:"lastname,omitempty"`
}

var client *mongo.Client

func createPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// marshal json
	var person Person
	json.NewDecoder(request.Body).Decode(&person)
	defer request.Body.Close()
	collection := client.Database("people").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, _ := collection.InsertOne(ctx, person)
	json.NewEncoder(response).Encode(result)
}

func getPeopleEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	// MongoDb uses cursors
	var people []Person
	collection := client.Database("people").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	cursor, err := collection.Find(ctx, bson.M{}) // bson.M{} - return everything from collection
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var person Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	// check if errors with cursor
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	// if all ok
	json.NewEncoder(response).Encode(people)
}

func getPersonEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var person Person
	collection := client.Database("people").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `"}`))
		return
	}
	json.NewEncoder(response).Encode(person)

	defer request.Body.Close()
}

func main() {
	fmt.Println("Starting application..")

	// MongoDb credentials
	var cred options.Credential
	cred.Username = "root"
	cred.Password = "example"

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// connect to MongoDb using credentials
	clientOption := options.Client().ApplyURI("mongodb://localhost:27017").SetAuth(cred)
	client, _ = mongo.Connect(ctx, clientOption)

	router := mux.NewRouter()
	router.HandleFunc("/person", createPersonEndpoint).Methods("POST")
	router.HandleFunc("/people", getPeopleEndpoint).Methods("GET")
	router.HandleFunc("/person/{id}", getPersonEndpoint).Methods("GET")

	http.ListenAndServe(":12345", router)

}
