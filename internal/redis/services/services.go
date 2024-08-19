package services

import (
	interface17 "conn/internal/interface"
	"conn/internal/models"
	"context"
)

type Services struct {
	O interface17.Origin
}

func (u *Services) Register(ctx context.Context, req *models.Register) error {
	return u.O.Register(ctx, req)
}

func (u *Services) Verify(ctx context.Context, req *models.Verify) error {
	return u.O.Verify(ctx, req)
}

func (u *Services) LogIn(ctx context.Context, req *models.LogIn) (string, error) {
	return u.O.LogIn(ctx, req)
}

func (u *Services) OriginAdd(ctx context.Context, req *models.OriginCreate) (string, error) {
	return u.O.OriginAdd(ctx, req)
}

func (u *Services) OriginGetbyId(ctx context.Context, req string) (*models.OriginCreate, error) {
	return u.O.OriginGetbyId(ctx, req)
}

func (u *Services) OriginGetAll(ctx context.Context) ([]*models.OriginCreate, error) {
	return u.O.OriginGetAll(ctx)
}

func (u *Services) OriginPut(ctx context.Context, req *models.OriginGet) error {
	return u.O.OriginPut(ctx, req)
}

func (u *Services) OriginDelete(ctx context.Context, req string) error {
	return u.O.OriginDelete(ctx, req)
}
