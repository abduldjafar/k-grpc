package repository

import "k-grpc/entitypb"

type Repository interface {
	Save(params ...interface{}) (string, error)
	GetByID(params ...interface{}) (*entitypb.TalentResponse, error)
	Update(data *entitypb.TalentUpdaterequest) error
	Delete(id string) error
	GetAll(params ...interface{}) (*entitypb.ListTalentsResponses, error)
}
