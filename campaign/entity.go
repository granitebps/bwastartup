package campaign

import (
	"github.com/granitebps/bwastartup/user"
	"github.com/leekchan/accounting"
	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	UserID           uint
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string

	CampaignImages []CampaignImage
	User           user.User
}

type CampaignImage struct {
	gorm.Model
	CampaignID uint
	FileName   string
	IsPrimary  int
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{
		Symbol:    "Rp",
		Precision: 2,
		Thousand:  ".",
		Decimal:   ",",
	}

	return ac.FormatMoney(c.GoalAmount)
}
