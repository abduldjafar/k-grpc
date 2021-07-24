package repository

import (
	"context"
	"k-grpc/config"
	"k-grpc/entitypb"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct{}

var (
	db = config.MongoDB()
)

func (*mongoRepository) Save(params ...interface{}) (string, error) {
	data := params[0]

	datas, err := db.Collection("users").InsertOne(context.TODO(), data)
	if err != nil {
		return "", err
	}

	id := datas.InsertedID.(primitive.ObjectID).Hex()

	return id, nil
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

func (*mongoRepository) Update(data *entitypb.TalentUpdaterequest) error {

	id := data.Id.Id

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	_, err = db.Collection("users").ReplaceOne(context.TODO(), bson.M{"_id": objectId}, data.RequestsData)

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (*mongoRepository) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	_, err = db.Collection("users").DeleteOne(context.TODO(), bson.M{"_id": objectId})

	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (*mongoRepository) GetAll(params ...interface{}) (*entitypb.ListTalentsResponses, error) {
	limit := params[0].(int64)
	index := params[1].(int64)
	results := entitypb.ListTalentsResponses{}

	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip(limit * index)

	cursor, err := db.Collection("users").Find(context.TODO(), bson.M{}, findOptions)

	if err != nil {
		log.Println(err.Error())
	}

	for cursor.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		elem := entitypb.TalentResponse{}
		err := cursor.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results.ListTalents = append(results.ListTalents, &elem)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &results, nil
}
func NewMongoRepository() Repository {
	return &mongoRepository{}
}
