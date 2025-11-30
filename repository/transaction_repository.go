package repository

import (
	"context"
	"go-mayar-payment-webhook/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	TransactionRepository interface {
		GetProductByID(ctx context.Context, tx *gorm.DB, productId uuid.UUID) (entity.Product, error)
		CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) (entity.Transaction, error)
	}

	transactionRepository struct {
		db *gorm.DB
	}
)

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db: db}
}

func (r *transactionRepository) GetProductByID(ctx context.Context, tx *gorm.DB, productId uuid.UUID) (entity.Product, error) {
	if tx == nil {
		tx = r.db
	}

	var product entity.Product
	if err := tx.WithContext(ctx).Where("id = ?", productId).First(&product).Error; err != nil {
		return entity.Product{}, err
	}

	return product, nil
}

func (r *transactionRepository) CreateTransaction(ctx context.Context, tx *gorm.DB, transaction entity.Transaction) (entity.Transaction, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&transaction).Error; err != nil {
		return entity.Transaction{}, err
	}

	return transaction, nil
}
