package repository

import (
	"context"

	"github.com/BacataCode/SmartCardConnectApiRest/models"
)

func InsertUser(ctx context.Context, user *models.InsertUser) (*models.Profile, error) {
	return implementation.InsertUser(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.Profile, error) {
	return implementation.GetUserById(ctx, id)
}

func ListUsers(ctx context.Context) ([]models.Profile, error) {
	return implementation.ListUsers(ctx)
}

func GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	return implementation.GetUserByEmail(ctx, email)
}

func UpdateUser(ctx context.Context, data models.UpdateUser) (*models.Profile, error) {
	return implementation.UpdateUser(ctx, data)
}

func DeleteUser(ctx context.Context, id string) error {
	return implementation.DeleteUser(ctx, id)
}

func UpdateUserPassword(ctx context.Context, userId string, newPassword string) (profile *models.Profile, err error) {
	return implementation.UpdateUserPassword(ctx, userId, newPassword)
}
