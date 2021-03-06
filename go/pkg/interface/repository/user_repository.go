package repository

import (
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/domain"
	"github.com/yawn-yawn-yawn/sampleNoteApp/go/pkg/usecase"
)

type userRepository struct {
	SQLHandler SQLHandler
}

func NewUserRepository(sh SQLHandler) usecase.UserRepository {
	return &userRepository{
		SQLHandler: sh,
	}
}

func (ur *userRepository) FindUserByUserID(userID string) (user domain.User, err error) {
	row := ur.SQLHandler.QueryRow("SELECT id, name, mail, created_at FROM users WHERE id=?", userID)
	if err = row.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
		return user, domain.InternalServerError(err)
	}
	return
}

func (ur *userRepository) FindUsers() (users domain.Users, err error) {
	rows, err := ur.SQLHandler.Query("SELECT id, name, mail, created_at FROM users")
	if err != nil {
		return users, domain.InternalServerError(err)
	}
	for rows.Next() {
		var user domain.User
		if err = rows.Scan(&user.ID, &user.Name, &user.Mail, &user.CreatedAt); err != nil {
			continue
		}
		users = append(users, user)
	}
	return
}
