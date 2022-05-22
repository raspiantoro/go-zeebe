package repository

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/approval/internal/model"
)

var _ ApprovalRepository = &approvalRepository{}

type approvalRepository struct {
	Option
}

func NewApprovalRepository(opt Option) *approvalRepository {
	return &approvalRepository{opt}
}

func (p *approvalRepository) Create(ctx context.Context, approval *model.Approval) (err error) {
	result := p.DB.Create(approval)
	if result.Error != nil {
		err = result.Error
		log.Error(err)
	}
	return
}
