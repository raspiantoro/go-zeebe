package commons

import (
	"github.com/raspiantoro/go-zeebe/approval/internal/command"
	"gorm.io/gorm"
)

type Option struct {
	Command command.Command
	DB      *gorm.DB
}
