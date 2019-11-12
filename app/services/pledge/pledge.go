package pledge

import (
	"github.com/jinzhu/gorm"
)

type PledgeRequest struct {
	UserID    uint `json:"user_id" validate:"required"`
	Amount    int  `json:"amount" validate:"required"`
	ProjectId uint `json:"project_id" validate:"required"`
}

func Create(r PledgeRequest, db *gorm.DB) {
}
