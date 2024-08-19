package interface17

import (
	"conn/internal/models"
	"context"
)

type Origin interface {
	Register(ctx context.Context, req *models.Register) error
	Verify(ctx context.Context, req *models.Verify) error
	LogIn(ctx context.Context, req *models.LogIn) (string,error)
	OriginAdd(ctx context.Context, req *models.OriginCreate) (string,error)
	OriginGetbyId(ctx context.Context, req string) (*models.OriginCreate, error)
	OriginGetAll(ctx context.Context) ([]*models.OriginCreate, error)
	OriginPut(ctx context.Context, req *models.OriginGet) (error)
	OriginDelete(ctx context.Context, req string) error
}
