package main

import (
	"context"

	"github.com/MWT-proger/go-keeper/configs/server/config"
	"github.com/MWT-proger/go-keeper/internal/pkg/logger"
	"github.com/MWT-proger/go-keeper/internal/server/handlers"
	"github.com/MWT-proger/go-keeper/internal/server/models"
	"github.com/MWT-proger/go-keeper/internal/server/rest"
	"github.com/MWT-proger/go-keeper/internal/server/services"
	"github.com/MWT-proger/go-keeper/internal/server/store"
)

type UserStore struct {
	*store.Store
}

func NewUserStore(baseStorage *store.Store) (*UserStore, error) {

	return &UserStore{baseStorage}, nil
}

func (s *UserStore) Insert(ctx context.Context, obj *models.User) error {
	return nil
}
func (s *UserStore) GetFirstByParameters(ctx context.Context, args map[string]interface{}) (*models.User, error) {
	user, _ := models.NewUser()
	return user, nil
}

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := run(ctx); err != nil {
		cancel()
		panic(err)
	}
}

// run() выполняет все предворительные действия и вызывает функцию запуска сервера
func run(ctx context.Context) error {
	// Просто заглушка для контекста пока
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var conf = config.InitConfig()

	if err := logger.Initialize(conf.LogLevel); err != nil {
		return err
	}
	storage, err := store.NewStore(ctx, conf)
	if err != nil {
		return err
	}

	userStorage, err := NewUserStore(storage)

	if err != nil {
		return err
	}

	userService := services.NewUserService(userStorage, conf)

	h, err := handlers.NewAPIHandler(userService)

	if err != nil {
		return err
	}

	err = rest.Run(h, conf)

	if err != nil {
		return err
	}

	return nil
}
