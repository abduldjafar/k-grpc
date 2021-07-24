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

func (*grpcTalent) AddTalent(ctx context.Context, request *entitypb.TalentRequest) (*entitypb.ID, error) {

	id, err := repo.Save(request)
	if err != nil {
		return nil, err
	}

	response := &entitypb.ID{
		Id: id,
	}

	return response, nil
}

func (*grpcTalent) GetTalent(ctx context.Context, request *entitypb.ID) (*entitypb.TalentResponse, error) {
	id := request.Id

	data, err := repo.GetByID(id)

	if err != nil {
		return &entitypb.TalentResponse{}, err
	}

	data.Id = request.Id
	return data, nil
}

func (*grpcTalent) UpdateTalent(ctx context.Context, request *entitypb.TalentUpdaterequest) (*entitypb.SuccessResponses, error) {

	if err := repo.Update(request); err != nil {
		return nil, err
	}

	return &entitypb.SuccessResponses{
		Message: "success",
	}, nil
}
func (*grpcTalent) DeleteTalent(ctx context.Context, request *entitypb.ID) (*entitypb.SuccessResponses, error) {
	id := request.GetId()

	if err := repo.Delete(id); err != nil {
		return nil, err
	}

	return &entitypb.SuccessResponses{
		Message: "success",
	}, nil
}
func (*grpcTalent) GetListTalents(ctx context.Context, request *entitypb.Pagination) (*entitypb.ListTalentsResponses, error) {
	datas, err := repo.GetAll(request.Limit, request.Page)

	if err != nil {
		return nil, err
	}
	return datas, nil
}
func NewTalentService() entitypb.TalentServiceServer {
	return &grpcTalent{}
}
