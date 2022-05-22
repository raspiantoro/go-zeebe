package service

import (
	"context"

	"github.com/raspiantoro/go-zeebe/approval/internal/commons"
	"github.com/raspiantoro/go-zeebe/approval/internal/payload"
	"github.com/raspiantoro/go-zeebe/approval/internal/repository"
)

type Option struct {
	commons.Option
	Repository repository.Repository
}

type Service struct {
	Approval ApprovalService
}

type ApprovalService interface {
	Submit(ctx context.Context, req *payload.ApprovalRequest) (purchaseID string, err error)
}
