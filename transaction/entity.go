package transaction

import (
	"github.com/granitebps/bwastartup/campaign"
	"github.com/granitebps/bwastartup/user"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Amount     int
	Status     string
	Code       string

	User     user.User
	Campaign campaign.Campaign
}
