package main

import (
	"context"
	"log"
	"os"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load environemnt variables
	err := godotenv.Load(".env")

	if err != nil {

		log.Fatal("Unable to load env: ", err)
	}

	log.Println("env loaded")

	// connect to db
	mongoClient, err = mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal("Unable to connect to database", err)
	}

	log.Println("connected to database")

	err = mongoClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("database ping error", err)
	}

	log.Println("ping to database successful")
}

func main() {
	// close the mongo connection
	defer mongoClient.Disconnect(context.Background())

	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", healthcheckHandler)

	log.Println("server is running")
	http.ListenAndServe(":5009", router)

}

func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("running..."))
}
