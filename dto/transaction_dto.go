package dto

import "errors"

const (
	MESSAGE_FAILED_GET_DATA_FROM_BODY  = "failed to get data from body"
	MESSAGE_FAILED_CREATE_TRANSACTION  = "failed to create transaction"
	MESSAGE_SUCCESS_CREATE_TRANSACTION = "success create transaction"
)

var (
	ErrFailedCreateInvoice = errors.New("failed to create invoice")
)

type (
	MayarItem struct {
		Quantity    int    `json:"quantity"`
		Rate        int    `json:"rate"` // price
		Description string `json:"description"`
	}

	MayarInvoice struct {
		Name         string      `json:"name"`
		Email        string      `json:"email"`
		MobileNumber string      `json:"mobile"`
		RedirectUrl  string      `json:"redirectUrl"`
		Description  string      `json:"description"`
		ExpiredAt    string      `json:"expiredAt"` // (format ISO 8601 date-time) 2025-12-01T09:41:09.401Z
		Items        []MayarItem `json:"items"`
	}

	CreateTransactionRequest struct {
		Name         string `json:"name" binding:"required"`
		Email        string `json:"email" binding:"required,email"`
		MobileNumber string `json:"mobile_number" binding:"required"`
		ProductID    string `json:"product_id" binding:"required,uuid"`
	}

	CreateTransactionResponse struct {
		InvoiceURL string `json:"invoice_url"`
	}
)
