package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/granitebps/bwastartup/helper"
	"github.com/granitebps/bwastartup/transaction"
	"github.com/granitebps/bwastartup/user"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign transactions",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	transactions, err := h.service.GetTransactionsByCampaignID(input)
	if err != nil {
		response := helper.APIResponse(
			"Failed to get campaign transactions",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"Campaign Transactions",
		http.StatusOK,
		"success",
		transaction.FormatCampaignTransactions(transactions),
	)
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID

	transactions, err := h.service.GetTransactionsByUserID(int(userID))
	if err != nil {
		response := helper.APIResponse(
			"Failed to get user transactions",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"User Transactions",
		http.StatusOK,
		"success",
		transaction.FormatUserTransactions(transactions),
	)
	c.JSON(http.StatusOK, response)
}
