package service

import (
	"context"
	"go-mayar-payment-webhook/constants"
	"go-mayar-payment-webhook/dto"
	"go-mayar-payment-webhook/entity"
	"go-mayar-payment-webhook/repository"
	"go-mayar-payment-webhook/utils/payment"
	"log"
	"time"

	"github.com/google/uuid"
)

type (
	TransactionService interface {
		CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error)
	}

	transactionService struct {
		transactionRepo repository.TransactionRepository
	}
)

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	return &transactionService{
		transactionRepo: transactionRepo,
	}
}

func (s *transactionService) CreateTransaction(ctx context.Context, req dto.CreateTransactionRequest) (dto.CreateTransactionResponse, error) {
	product, err := s.transactionRepo.GetProductByID(ctx, nil, uuid.MustParse(req.ProductID))
	if err != nil {
		return dto.CreateTransactionResponse{}, err
	}

	item := dto.MayarItem{
		Quantity:    1,
		Rate:        product.Price,
		Description: product.Name,
	}

	invoice := dto.MayarInvoice{
		Name:         req.Name,
		Email:        req.Email,
		MobileNumber: req.MobileNumber,
		RedirectUrl:  constants.REDIRECT_URL,
		Description:  "Payment for " + product.Name,
		ExpiredAt:    time.Now().Add(time.Hour * 24 * 30 * 6).Format("2006-01-02T15:04:05Z"),
		Items:        []dto.MayarItem{item},
	}

	res, err := payment.SendMayarInvoice(invoice)
	if err != nil {
		return dto.CreateTransactionResponse{}, dto.ErrFailedCreateInvoice
	}

	log.Println(res)

	invoiceURL, ok := res["data"].(map[string]interface{})["link"].(string)
	if !ok {
		log.Println("invoice_url tidak ditemukan atau bukan string")
		return dto.CreateTransactionResponse{}, err
	} else {
		log.Println("Invoice URL:", invoiceURL)
	}

	transID, ok := res["data"].(map[string]interface{})["id"].(string)
	if !ok {
		log.Println("transID tidak ditemukan atau bukan string")
		return dto.CreateTransactionResponse{}, err
	} else {
		log.Println("Transaction ID:", transID)
	}

	transaction := entity.Transaction{
		ID:         uuid.MustParse(transID),
		ProductID:  product.ID,
		AmountPaid: 0,
		Status:     "PENDING",
		InvoiceUrl: invoiceURL,
	}

	trans, err := s.transactionRepo.CreateTransaction(ctx, nil, transaction)
	if err != nil {
		return dto.CreateTransactionResponse{}, err
	}

	return dto.CreateTransactionResponse{
		InvoiceURL: trans.InvoiceUrl,
	}, nil
}
