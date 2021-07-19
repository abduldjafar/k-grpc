package service

import (
	"context"
	"k-grpc/entitypb"
	"k-grpc/repository"
	"log"
)

type grpcTalent struct{}

var (
	repo repository.Repository = repository.NewMongoRepository()
)

func (*grpcTalent) AddTalent(ctx context.Context, request *entitypb.TalentRequest) (*entitypb.TalentResponse, error) {
	data := map[string]interface{}{}

	data["age"] = request.Age
	data["email"] = request.Email
	data["name"] = request.Name

	if err := repo.Save(data); err != nil {
		return &entitypb.TalentResponse{}, err
	}

	response := &entitypb.TalentResponse{
		Email: request.Email,
	}

	return response, nil
}

func (*grpcTalent) GetTalent(ctx context.Context, request *entitypb.ID) (*entitypb.TalentResponse, error) {
	id := request.Id

	data, err := repo.GetByID(id)

	if err != nil {
		return &entitypb.TalentResponse{}, err
	}

	datas := data.(map[string]interface{})

	log.Println(datas)

	response := &entitypb.TalentResponse{
		Email: datas["email"].(string),
	}
	return response, nil
}

func NewTalentService() entitypb.TalentServiceServer {
	return &grpcTalent{}
}
