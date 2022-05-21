package service

import (
	"context"
	"errors"
	"fmt"
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

func (p *purchaseService) Prepare(ctx context.Context, variables map[string]interface{}) (purchaseID string, err error) {
	createdDate := time.Now()

	item, ok := variables["item"].(string)
	if !ok {
		err = errors.New("invalid item value")
		log.Error(err)
		return
	}

	fmt.Println("Price: ", variables["price"].(float64))
	price, ok := variables["price"].(float64)
	if !ok {
		err = errors.New("invalid price value")
		log.Error(err)
		return
	}

	processKey, ok := variables["process_key"].(int64)
	if !ok {
		err = errors.New("invalid process_key value")
		log.Error(err)
		return
	}

	model := model.Purchase{
		Item:       item,
		Price:      uint64(price),
		Status:     "waiting-approval",
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
