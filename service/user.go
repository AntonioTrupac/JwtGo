package service

import (
	"context"
	"strings"

	"antonio.trupac/config"
	"antonio.trupac/graph/model"
	"antonio.trupac/models"
	"antonio.trupac/tools"

	"github.com/google/uuid"
)

func UserCreate(ctx context.Context, input model.NewUser) (*models.User, error) {
	db := config.GetDB()

	input.Password = tools.HashPassword(input.Password)

	user := models.User{
		ID:       int(uuid.New().ID()),
		Name:     input.Name,
		Email:    strings.ToLower(input.Email),
		Password: input.Password,
	}

	if err := db.Model(user).Create(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserById(ctx context.Context, id int) (*models.User, error) {
	db := config.GetDB()

	var user models.User

	if err := db.Model(user).Where("id = ?", id).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	db := config.GetDB()

	var user models.User
	if err := db.Model(user).Where("email LIKE ?", email).Take(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
