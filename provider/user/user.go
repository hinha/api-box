package user

import (
	"context"
	"github.com/hinha/api-box/entity"
	"github.com/hinha/api-box/provider"
	"github.com/hinha/api-box/provider/user/usecase"
)

// User provide function for managing user data
type User struct {
	db provider.DB
}

// Fabricate user provider
func Fabricate(db provider.DB) *User {
	return &User{db: db}
}

func (u *User) CreateOAuth(ctx context.Context, userInsertable entity.GoogleUser) (int, *entity.ApplicationError) {
	create := usecase.Create{}
	return create.Perform(ctx, userInsertable, u.db)
}
