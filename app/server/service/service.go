package service

import (
	"crac55/model"
)

type Service struct {
	Dao *model.Dao
}

func NewService() *Service {
	return &Service{
		Dao: model.NewDao(),
	}
}
