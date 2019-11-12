package pledge

import (
	"github.com/jinzhu/gorm"
	"github.com/jmramos02/smarty-seed-backend/app/models"
)

type PledgeRequest struct {
	UserID    uint `json:"user_id" validate:"required"`
	Amount    int  `json:"amount" validate:"required"`
	ProjectID uint `json:"project_id" validate:"required"`
}

func Create(pledge models.Pledge, db *gorm.DB) models.Pledge {
	db.NewRecord(pledge)
	db.Set("gorm:association_autocreate", false)Create(&pledge)

	return pledge
}
