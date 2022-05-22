package repository

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/purchase/internal/model"
)

var _ PurchaseRepository = &purchaseRepository{}

type purchaseRepository struct {
	Option
}

func NewPurchaseRepository(opt Option) *purchaseRepository {
	return &purchaseRepository{opt}
}

func (p *purchaseRepository) Get(ctx context.Context, ID string) (purchase model.Purchase, err error) {
	result := p.DB.Where("id = ?", ID).First(&purchase)
	if result.Error != nil {
		err = result.Error
		log.Error(err)
	}
	return
}

func (p *purchaseRepository) GetByStatus(ctx context.Context, status string) (purchase model.Purchase, err error) {
	result := p.DB.Where("status = ?", status).First(&purchase)
	if result.Error != nil {
		err = result.Error
		log.Error(err)
	}
	return
}

func (p *purchaseRepository) Create(ctx context.Context, purchase *model.Purchase) (err error) {
	result := p.DB.Create(purchase)
	if result.Error != nil {
		err = result.Error
		log.Error(err)
	}
	return
}

func (p *purchaseRepository) Update(ctx context.Context, purchase *model.Purchase) (err error) {
	result := p.DB.Save(purchase)
	if result.Error != nil {
		err = result.Error
		log.Error(err)
	}
	return
}
