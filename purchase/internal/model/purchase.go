package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Item       string    `json:"item"`
	Price      uint64    `json:"price"`
	Status     string    `json:"status"`
	ProcessKey int64     `json:"process_key"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Purchase) TableName() string {
	return "purchase"
}

func (p *Purchase) BeforeCreate(tx *gorm.DB) error {
	p.ID = uuid.New()
	return nil
}
