package transaction

import "crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}
