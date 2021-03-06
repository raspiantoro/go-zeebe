package service

import (
	"context"

	"github.com/raspiantoro/go-zeebe/purchase/internal/commons"
	"github.com/raspiantoro/go-zeebe/purchase/internal/repository"
)

type Option struct {
	commons.Option
	Repository repository.Repository
}

type Service struct {
	Purchase PurchaseService
}

type PurchaseService interface {
	Prepare(ctx context.Context, status string, variables map[string]interface{}) (purchaseID string, err error)
	UpdateStatus(ctx context.Context, status string, variables map[string]interface{}) (err error)
}
