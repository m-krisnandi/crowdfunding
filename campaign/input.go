package campaign

type GetCampaignDetailInput struct {
	ID string `uri:"id" binding:"required"`
}
