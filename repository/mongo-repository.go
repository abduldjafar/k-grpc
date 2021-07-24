package repository

import (
	"context"
	"k-grpc/config"
	"k-grpc/entitypb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type mongoRepository struct{}

var (
	db = config.MongoDB()
)

func (*mongoRepository) Save(params ...interface{}) error {
	data := params[0]

	_, err := db.Collection("users").InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}

	return nil
}

func (*mongoRepository) GetByID(params ...interface{}) (*entitypb.TalentResponse, error) {
	var data *entitypb.TalentResponse

	id := params[0].(string)

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	result := db.Collection("users").FindOne(context.TODO(), bson.M{"_id": objectId})

	if err := result.Decode(&data); err != nil {
		return nil, err
	}

	return data, nil
}

func NewMongoRepository() Repository {
	return &mongoRepository{}
}
