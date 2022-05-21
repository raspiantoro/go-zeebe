package repository

import (
	"context"

	"github.com/raspiantoro/go-zeebe/purchase/internal/commons"
	"github.com/raspiantoro/go-zeebe/purchase/internal/model"
)

type Option struct {
	commons.Option
}

type Repository struct {
	Purchase PurchaseRepository
}

type PurchaseRepository interface {
	Get(ctx context.Context, ID string) (purchase model.Purchase, err error)
	GetByStatus(ctx context.Context, status string) (purchase model.Purchase, err error)
	Create(ctx context.Context, purchase *model.Purchase) (err error)
	Update(ctx context.Context, purchase *model.Purchase) (err error)
}
