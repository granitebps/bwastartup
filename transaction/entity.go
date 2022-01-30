package transaction

import (
	"github.com/granitebps/bwastartup/campaign"
	"github.com/granitebps/bwastartup/user"
	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	CampaignID uint
	UserID     uint
	Amount     int
	Status     string
	Code       string
	PaymentURL string

	User     user.User
	Campaign campaign.Campaign
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
		Thousand:  ".",
		Decimal:   ",",
	}

	return ac.FormatMoney(t.Amount)
}
