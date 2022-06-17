package repositories

import (
	"api-bootstrap-echo/models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func ConnectDB() *mongo.Client {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017/mobo?readPreference=primary&appname=MongoDB%20Compass&ssl=false"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	//ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB")
	return client
}

//Client instance
var DB *mongo.Client = ConnectDB()

// NewBootstrapRepository :
func NewBootstrapRepository() *BootstrapRepository {
	return &BootstrapRepository{}
}

// BootstrapRepository :
type BootstrapRepository struct {
}

// Get :
func (BootstrapRepository) Get(id int) (models.Bootstrap, error) {
	bootstrap := models.Bootstrap{ID: id, Title: "Ingeniero en Sistemas", Description: "Se busca ingeniero en sistemas"}
	return bootstrap, nil
}

// Save :
func (BootstrapRepository) Save(bootstrap models.Bootstrap) (bool, error) {
	return true, nil
}

// go get -u go.mongodb.org/mongo-driver/mongo github.com/go-playground/validator/v10
