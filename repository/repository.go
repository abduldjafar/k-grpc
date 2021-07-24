package repository

import "k-grpc/entitypb"

type Repository interface {
	Save(params ...interface{}) error
	GetByID(params ...interface{}) (*entitypb.TalentResponse, error)
}
