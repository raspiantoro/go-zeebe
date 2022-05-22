package service

import (
	"context"
	"errors"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/raspiantoro/go-zeebe/purchase/internal/model"
)

type purchaseService struct {
	Option
}

// makesure purchaseService implementing PurchaseService interface
var _ PurchaseService = &purchaseService{}

func NewPurchaseService(opt Option) *purchaseService {
	return &purchaseService{opt}
}

func (p *purchaseService) Prepare(ctx context.Context, status string, variables map[string]interface{}) (purchaseID string, err error) {
	createdDate := time.Now()

	item, ok := variables["item"].(string)
	if !ok {
		err = errors.New("invalid item")
		log.Error(err)
		return
	}

	price, ok := variables["price"].(float64)
	if !ok {
		err = errors.New("invalid price")
		log.Error(err)
		return
	}

	processKey, ok := variables["process_key"].(int64)
	if !ok {
		err = errors.New("invalid process_key")
		log.Error(err)
		return
	}

	model := model.Purchase{
		Item:       item,
		Price:      uint64(price),
		Status:     status,
		ProcessKey: processKey,
		CreatedAt:  createdDate,
		UpdatedAt:  createdDate,
	}

	err = p.Repository.Purchase.Create(ctx, &model)
	if err != nil {
		return
	}

	purchaseID = model.ID.String()

	return
}

func (p *purchaseService) UpdateStatus(ctx context.Context, status string, variables map[string]interface{}) (err error) {
	purchaseID, ok := variables["purhcase_id"].(string)
	if !ok {
		err = errors.New("invalid purhcase_id")
		log.Error(err)
	}

	purchase, err := p.Repository.Purchase.Get(ctx, purchaseID)
	if err != nil {
		return
	}

	purchase.Status = status

	purchase.UpdatedAt = time.Now()

	err = p.Repository.Purchase.Update(ctx, &purchase)

	return
}
