package controller

import (
	"fmt"
	"note-app/usecase"
)

type userController struct {
	UserInteractor usecase.UserInteractor
}

type UserController interface {
	Create(*CreateAccountRequest) (*CreateAccountResponse, error)
	Delete(string) (*DeleteUserResponse, error)
	Users() (*GetUsersResponse, error)
	UserByID(string) (*GetUserResponse, error)
}

func NewUserController(ui usecase.UserInteractor) UserController {
	return &userController{
		UserInteractor: ui,
	}
}

func (uc *userController) Create(req *CreateAccountRequest) (*CreateAccountResponse, error) {
	user, err := uc.UserInteractor.Add(req.ID, req.Name, req.Password, req.Mail)
	if err != nil {
		return &CreateAccountResponse{}, err
	}
	return &CreateAccountResponse{
		ID:        user.ID,
		Name:      user.Name,
		Password:  user.Password,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt,
	}, nil

}

type CreateAccountRequest struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

type CreateAccountResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}

func (uc *userController) Delete(userID string) (*DeleteUserResponse, error) {
	err := uc.UserInteractor.Delete(userID)
	if err != nil {
		return &DeleteUserResponse{}, err
	}
	return &DeleteUserResponse{
		Message: fmt.Sprintf("success: delete %s", userID),
	}, nil
}

type DeleteUserResponse struct {
	Message string `json:"message"`
}

func (uc *userController) Users() (*GetUsersResponse, error) {
	users, err := uc.UserInteractor.ShowUsers()
	if err != nil {
		return &GetUsersResponse{}, err
	}
	var response GetUsersResponse
	for _, v := range users {
		var res GetUserResponse
		res.ID = v.ID
		res.Name = v.Name
		res.Mail = v.Mail
		res.CreatedAt = v.CreatedAt
		response.Users = append(response.Users, res)
	}
	return &response, nil
}

type GetUsersResponse struct {
	Users []GetUserResponse `json:"users"`
}

func (uc *userController) UserByID(userID string) (*GetUserResponse, error) {
	user, err := uc.UserInteractor.ShowUserByID(userID)
	if err != nil {
		return &GetUserResponse{}, err
	}
	return &GetUserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Mail:      user.Mail,
		CreatedAt: user.CreatedAt,
	}, nil
}

type GetUserResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Mail      string `json:"mail"`
	CreatedAt string `json:"created_at"`
}
