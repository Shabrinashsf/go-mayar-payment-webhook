package controller

import (
	"go-mayar-payment-webhook/dto"
	"go-mayar-payment-webhook/service"
	"go-mayar-payment-webhook/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

type (
	TransactionController interface {
		CreateTransaction(ctx *gin.Context)
	}

	transactionController struct {
		transactionService service.TransactionService
	}
)

func NewTransactionController(transactionService service.TransactionService) TransactionController {
	return &transactionController{
		transactionService: transactionService,
	}
}

func (c *transactionController) CreateTransaction(ctx *gin.Context) {
	// Dont forget to implement your business logic
	var req dto.CreateTransactionRequest
	if err := ctx.ShouldBind(&req); err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_GET_DATA_FROM_BODY, err.Error(), nil)
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	result, err := c.transactionService.CreateTransaction(ctx.Request.Context(), req)
	if err != nil {
		res := response.BuildResponseFailed(dto.MESSAGE_FAILED_CREATE_TRANSACTION, err.Error(), nil)
		ctx.JSON(http.StatusInternalServerError, res)
		return
	}

	res := response.BuildResponseSuccess(dto.MESSAGE_SUCCESS_CREATE_TRANSACTION, result)
	ctx.JSON(http.StatusOK, res)
}
