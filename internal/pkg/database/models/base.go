package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Base struct {
	ID        string `json:"id" gorm:"primary_key" default:uuid_generate_v4()`
	UpdatedAt int64  `gorm:"autoUpdateTime:milli"`
	CreatedAt int64  `gorm:"autoUpdateTime:milli"`
}

func (b *Base) BeforeCreate(_ *gorm.DB) error {
	if b.ID == "" {
		b.ID = uuid.NewString()
	}

	return nil
}
