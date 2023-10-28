package transaction

import "crowdfunding/user"

type GetCampaignTransactionsInput struct {
	ID   string `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	CampaignID string `json:"campaign_id" binding:"required"`
	Amount     int    `json:"amount" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
}
