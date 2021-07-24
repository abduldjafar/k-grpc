package service

import (
	"context"
	"k-grpc/entitypb"
	"k-grpc/repository"
)

type grpcTalent struct{}

var (
	repo repository.Repository = repository.NewMongoRepository()
)

func (*grpcTalent) AddTalent(ctx context.Context, request *entitypb.TalentRequest) (*entitypb.SuccessResponses, error) {

	if err := repo.Save(request); err != nil {
		return nil, err
	}

	response := &entitypb.SuccessResponses{
		Message: "success",
	}

	return response, nil
}

func (*grpcTalent) GetTalent(ctx context.Context, request *entitypb.ID) (*entitypb.TalentResponse, error) {
	id := request.Id

	data, err := repo.GetByID(id)

	if err != nil {
		return &entitypb.TalentResponse{}, err
	}

	return data, nil
}

func (*grpcTalent) UpdateTalent(ctx context.Context, request *entitypb.TalentRequest) (*entitypb.SuccessResponses, error) {
	return &entitypb.SuccessResponses{}, nil
}
func (*grpcTalent) DeleteTalent(ctx context.Context, request *entitypb.TalentRequest) (*entitypb.SuccessResponses, error) {
	return &entitypb.SuccessResponses{}, nil
}
func (*grpcTalent) GetListTalents(ctx context.Context, request *entitypb.Pagination) (*entitypb.ListTalentsResponses, error) {
	return &entitypb.ListTalentsResponses{}, nil
}
func NewTalentService() entitypb.TalentServiceServer {
	return &grpcTalent{}
}
