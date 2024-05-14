package auth

import (
	"context"
	"testing"

	"github.com/google/uuid"
)

func TestWithUserID(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()

	ctx = WithUserID(ctx, userID)

	v, ok := ctx.Value(userIDKey).(uuid.UUID)
	if !ok {
		t.Errorf("ожидаемый идентификатор пользователя должен быть типа uuid.UUID")
	}

	if v != userID {
		t.Errorf("ожидаемый идентификатор пользователя должен быть %v, получил %v", userID, v)
	}
}

func TestUserIDFrom(t *testing.T) {
	ctx := context.Background()
	userID := uuid.New()

	ctx = context.WithValue(ctx, userIDKey, userID)

	v, ok := UserIDFrom(ctx)
	if !ok {
		t.Errorf("ожидаемый идентификатор пользователя должен присутствовать в контексте")
	}

	if v != userID {
		t.Errorf("ожидаемый идентификатор пользователя должен быть %v, получил %v", userID, v)
	}
}
