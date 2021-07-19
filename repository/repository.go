package repository

type Repository interface {
	Save(params ...interface{}) error
	GetByID(params ...interface{}) (interface{}, error)
}
