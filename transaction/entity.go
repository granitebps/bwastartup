package transaction

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
}
