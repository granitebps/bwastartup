package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/granitebps/bwastartup/campaign"
	"github.com/granitebps/bwastartup/helper"
)

type campaignHandler struct {
	service campaign.Service
}

func NewCampaignHandler(service campaign.Service) *campaignHandler {
	return &campaignHandler{service}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.service.GetCampaigns(userId)
	if err != nil {
		response := helper.APIResponse(
			"Error to get campaigns",
			http.StatusBadRequest,
			"error",
			nil,
		)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse(
		"List of campaigns",
		http.StatusOK,
		"success",
		campaign.FormatCampaigns(campaigns),
	)
	c.JSON(http.StatusOK, response)
}
