package service

import (
	"crac55/app/server/model"
)

type Service struct {
	Dao *model.Dao
}

func NewService() *Service {
	return &Service{
		Dao: model.NewDao(),
	}
}
