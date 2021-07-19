package config

import (
	"context"
	"log"
	"os"

	"github.com/BurntSushi/toml"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Configuration struct {
	MongoDb mongodb
}

type mongodb struct {
	Url      string
	Ip       string
	Port     string
	Database string
	User     string
	Password string
	Ssl      string
}

func GetConfig(baseConfig *Configuration) {
	basePath, _ := os.Getwd()
	if _, err := toml.DecodeFile(basePath+"/config.toml", &baseConfig); err != nil {
		log.Println(err)
	}
}

func MongoDB() *mongo.Database {
	baseConfig := &Configuration{}
	GetConfig(baseConfig)

	clientOptions := options.Client().ApplyURI("mongodb+srv://" + baseConfig.MongoDb.User + ":" + baseConfig.MongoDb.Password + "@" + baseConfig.MongoDb.Url + "/myFirstDatabase?retryWrites=true&w=majority")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to MongoDB!")

	database := client.Database(baseConfig.MongoDb.Database)
	return database

}
