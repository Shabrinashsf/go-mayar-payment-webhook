package routes

import (
	"go-mayar-payment-webhook/controller"

	"github.com/gin-gonic/gin"
)

func Transaction(route *gin.Engine, transactionController controller.TransactionController) {
	routes := route.Group("/transaction")
	{
		routes.POST("/buy", transactionController.CreateTransaction)
	}
}
