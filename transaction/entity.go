package transaction

import (
	"crowdfunding/user"
	"time"
)

type Transaction struct {
	ID         string
	CampaignID string
	UserID     string
	Amount     int
	Status     string
	Code       string
	User       user.User
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
