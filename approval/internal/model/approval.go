package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Approval struct {
	ID         uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	PurchaseID uuid.UUID `json:"purchase_id"`
	Username   string    `json:"username"`
	Action     string    `json:"action"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (Approval) TableName() string {
	return "approval"
}

func (a *Approval) BeforeCreate(tx *gorm.DB) error {
	a.ID = uuid.New()
	return nil
}
