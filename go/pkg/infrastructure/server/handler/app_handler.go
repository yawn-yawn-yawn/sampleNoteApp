package handler

import (
	"errors"
	"net/http"

	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/middleware"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/infrastructure/server/response"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/interface/repository"
)

type appHandler struct {
	AccountHandler AccountHandler
	UserHandler    UserHandler
	NoteHandler    NoteHandler
}

type AppHandler interface {
	// account
	ManageAccount() http.HandlerFunc
	Login() http.HandlerFunc

	// user
	GetUserByUserID() http.HandlerFunc
	GetUsers() http.HandlerFunc

	// note
	ManageANote() http.HandlerFunc
	ManageNotes() http.HandlerFunc
}

func NewAppHandler(sh repository.SQLHandler) AppHandler {
	return &appHandler{
		AccountHandler: NewAccountHandler(sh),
		UserHandler:    NewUserHandler(sh),
		NoteHandler:    NewNoteHandler(sh),
	}
}

func (ah *appHandler) ManageAccount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.AccountHandler.GetAccount).ServeHTTP(w, r)
		case http.MethodPost:
			ah.AccountHandler.CreateAccount(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.AccountHandler.UpdateAccount).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.AccountHandler.DeleteAccount).ServeHTTP(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			ah.AccountHandler.Login(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) GetUserByUserID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.UserHandler.GetUserByUserID(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) GetUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			ah.UserHandler.GetUsers(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageANote() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.NoteHandler.GetNoteByNoteID).ServeHTTP(w, r)
		case http.MethodPut:
			middleware.Authorized(ah.NoteHandler.UpdateNote).ServeHTTP(w, r)
		case http.MethodDelete:
			middleware.Authorized(ah.NoteHandler.DeleteNote).ServeHTTP(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}

func (ah *appHandler) ManageNotes() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			middleware.Authorized(ah.NoteHandler.GetNotes).ServeHTTP(w, r)
		case http.MethodPost:
			middleware.Authorized(ah.NoteHandler.CreateNote).ServeHTTP(w, r)
		default:
			response.HttpError(w, domain.MethodNotAllowed(errors.New("method not allowed")))
		}
	}
}
