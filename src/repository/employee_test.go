package repository

import (
	"context"
	"log"
	"os"
	"renda/employee/src/model"
	"testing"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func newMongoClient() *mongo.Client {
	mongoTestClient, err := mongo.Connect(context.Background(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

	if err != nil {
		log.Fatal("Error while connecting to mongoDB", err)
	}

	log.Println("Connected to mongoDB")

	err = mongoTestClient.Ping(context.Background(), readpref.Primary())

	if err != nil {
		log.Fatal("Ping failed", err)

	}
	log.Println("Ping successful")

	return mongoTestClient
}

func TestMongoOperations(t *testing.T) {
	err := godotenv.Load("../../.env.test")

	if err != nil {
		t.Fatal("Unable to load env", err)
	}

	t.Log("env loaded")

	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// dummy data
	emp1 := "TEST-001"

	// connect to collection
	coll := mongoTestClient.Database(os.Getenv("DATABASE_NAME")).Collection(os.Getenv("COLLECTION_NAME"))

	empRepo := EmployeeRepo{MongoCollection: coll}

	t.Run("Insert Employee", func(t *testing.T) {
		emp := model.Employee{
			FirstName:  "Maxwell",
			LastName:   "Ihiaso",
			Age:        50,
			Position:   "Software Developer",
			Salary:     100000,
			EmployeeID: emp1,
		}

		result, err := empRepo.InsertEmployee(&emp)

		if err != nil {
			t.Fatal("Insert Employee operation failed", err)
		}

		t.Log("Insert Employee operation passed", result)
	})
}
