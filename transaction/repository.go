package transaction

import "gorm.io/gorm"

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetByCampaignID(campaignID string) ([]Transaction, error)
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetByCampaignID(campaignID string) ([]Transaction, error) {
	var transaction []Transaction

	err := r.db.Where("campaign_id = ?", campaignID).Preload("User").Order("created_at desc").Find(&transaction).Error
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
