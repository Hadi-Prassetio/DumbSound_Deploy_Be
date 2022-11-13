package repositories

import (
	"dumbsound/models"

	"gorm.io/gorm"
)

type TransactionRepository interface {
	FindTransactions() ([]models.Transaction, error)
	GetTransaction(ID int) (models.Transaction, error)
	GetTransactionActice(UserID int) (models.Transaction, error)
	CreateTransaction(transaction models.Transaction) (models.Transaction, error)
	UpdateTransaction(status string, ID string) error
}

func RepositoryTransaction(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindTransactions() ([]models.Transaction, error) {
	var transacrions []models.Transaction

	err := r.db.Preload("User").Find(&transacrions).Error
	return transacrions, err
}

func (r *repository) GetTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction

	err := r.db.Preload("User").First(&transaction, "id = ?", ID).Error
	return transaction, err
}

func (r *repository) GetTransactionActice(UserID int) (models.Transaction, error) {
	var transacrion models.Transaction
	err := r.db.Preload("User").Where("user_id = ? AND status_user = ?", UserID, "Active").First(&transacrion).Error
	return transacrion, err
}

func (r *repository) CreateTransaction(transaction models.Transaction) (models.Transaction, error) {
	err := r.db.Create(&transaction).Error

	return transaction, err
}

// Create UpdateTransaction method here ...
func (r *repository) UpdateTransaction(status string, ID string) error {
	var transaction models.Transaction
	r.db.First(&transaction, ID)

	if status != transaction.Status && status == "success" {
		transaction.Status = status
		transaction.Limit = 30
		transaction.StatusUser = "Active"
	}
	if status != transaction.Status && status == "failed" {
		transaction.Status = status
		transaction.Limit = 0
		transaction.StatusUser = "Not Active"
	}

	err := r.db.Save(&transaction).Error

	return err
}
