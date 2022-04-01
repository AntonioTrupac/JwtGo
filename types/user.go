package types

import (
	"antonio.trupac/graph/model"
	"antonio.trupac/models"
)

func MapUser(payload *models.User) *model.User {
	return &model.User{
		ID:    payload.ID,
		Name:  payload.Name,
		Email: payload.Email,
	}
}
