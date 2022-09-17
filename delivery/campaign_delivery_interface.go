package delivery

import (
	"github.com/gin-gonic/gin"
)

type CampaignDelivery interface {
	GetCampaigns(c *gin.Context)
}
