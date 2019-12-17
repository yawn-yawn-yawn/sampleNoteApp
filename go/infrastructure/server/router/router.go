package router

import (
	"note-app/infrastructure/server"
	"note-app/infrastructure/server/handler"
)

// SetUpRouting ルーティング
func SetUpRouting(s *server.Server, h *handler.AppHandler) {
	s.Get("/users", h.UserHandler.GetUsers)
	s.Get("/users/", h.UserHandler.GetUserByID)
	s.Post("/account/create", h.UserHandler.CreateUser)
	s.Post("/account/delete/", h.UserHandler.DeleteUser)
}
