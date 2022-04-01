package service

import (
	"context"

	"antonio.trupac/graph/model"
	"antonio.trupac/tools"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func UserRegister(ctx context.Context, input model.NewUser) (interface{}, error) {
	// check the email
	_, err := GetUserByEmail(ctx, input.Email)

	if err == nil {
		if err != gorm.ErrRecordNotFound {
			return nil, err
		}
	}

	createdUser, err := UserCreate(ctx, input)

	if err != nil {
		return nil, err
	}

	// generate the token with the created user id
	token, err := GenerateJwt(ctx, createdUser.ID)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}

func UserLogin(ctx context.Context, email string, password string) (interface{}, error) {
	getUser, err := GetUserByEmail(ctx, email)

	if err != nil {
		// if user not found
		if err == gorm.ErrRecordNotFound {
			return nil, &gqlerror.Error{
				Message: "Email not found",
			}
		}

		return nil, err
	}

	if err := tools.ComparePassword(getUser.Password, password); err != nil {
		return nil, err
	}

	token, err := GenerateJwt(ctx, getUser.ID)

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"token": token,
	}, nil
}
