package repository

import (
	"context"

	"github.com/raspiantoro/go-zeebe/approval/internal/commons"
	"github.com/raspiantoro/go-zeebe/approval/internal/model"
)

type Option struct {
	commons.Option
}

type Repository struct {
	Approval ApprovalRepository
}

type ApprovalRepository interface {
	Create(ctx context.Context, purchase *model.Approval) (err error)
}
