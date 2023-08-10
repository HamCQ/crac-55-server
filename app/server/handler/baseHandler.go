package handler

import (
	"crac55/app/server/service"
)

type CracHandler struct {
	Service *service.Service
}

func NewCrac() *CracHandler {
	return &CracHandler{
		Service: service.NewService(),
	}
}
