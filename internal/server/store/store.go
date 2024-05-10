package store

import (
	"context"

	"github.com/MWT-proger/go-keeper/internal/server/models"
)

type UserStore struct{}

func NewUserStore() (*UserStore, error) {

	return &UserStore{}, nil
}

func (s *UserStore) Insert(ctx context.Context, obj *models.User) error {
	return nil
}
func (s *UserStore) GetFirstByParameters(ctx context.Context, args map[string]interface{}) (*models.User, error) {
	user, _ := models.NewUser()
	return user, nil
}
