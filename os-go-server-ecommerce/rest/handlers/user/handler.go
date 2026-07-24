package user

import (
	"ecommerce/config"
)

type Handler struct {
	cnf *config.Config
	// userRepo repo.UserRepo
	svc Service
}

func NewHandler(cnf *config.Config, svc Service) *Handler {
	return &Handler{
		cnf: cnf,
		svc: svc,
	}
}
