package service

import (
	"context"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/approval/internal/model"
	"github.com/raspiantoro/go-zeebe/approval/internal/payload"
)

type approvalService struct {
	Option
}

// makesure purchaseService implementing PurchaseService interface
var _ ApprovalService = &approvalService{}

func NewApprovalService(opt Option) *approvalService {
	return &approvalService{opt}
}

func (p *approvalService) Submit(ctx context.Context, req *payload.ApprovalRequest) (approvalID string, err error) {
	createdDate := time.Now()

	purchaseID, err := uuid.Parse(req.PurchaseID)
	if err != nil {
		log.Error(err)
	}

	model := model.Approval{
		PurchaseID: purchaseID,
		Username:   req.Username,
		Action:     strings.ToUpper(req.Action),
		CreatedAt:  createdDate,
		UpdatedAt:  createdDate,
	}

	err = p.Repository.Approval.Create(ctx, &model)
	if err != nil {
		return
	}

	approvalID = model.ID.String()

	return
}
