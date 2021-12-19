package campaign

import (
	"github.com/granitebps/bwastartup/user"
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
