package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	pb "workout3/pb"

	"google.golang.org/grpc"
)

type Person struct {
	Name string `json:"name"`
}

func someData(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, "error in decoding", http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		http.Error(w, "failed to connect to gRPC ", http.StatusInternalServerError)
		return
	}
	defer conn.Close()

	client := pb.NewMockServiceClient(conn)
	res, err := client.GetSomeData(context.Background(), &pb.UserData{Name: person.Name})
	if err != nil {
		http.Error(w, "failed to get response from gRPC ", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(res)
}

func main() {
	http.HandleFunc("/name", someData)

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Println("error listen in port 8081")
	}
}
