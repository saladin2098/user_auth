package handlers

import (
	"github.com/Mubinabd/auth_service/service"
)

type Handler struct {
	UserService *service.UserService
}

func NewHandler(userService *service.UserService) *Handler {
	return &Handler{
        UserService: userService,
    }
}
