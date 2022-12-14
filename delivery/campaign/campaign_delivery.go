package campaigndelivery

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	campaignusecase "service-campaign-startup/usecase/campaign"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
)

type campaignDelivery struct {
	campaignUseCase campaignusecase.CampaignUseCase
}

func NewCampaignDelivery(campaignUseCase campaignusecase.CampaignUseCase) CampaignDelivery {
	return &campaignDelivery{
		campaignUseCase: campaignUseCase,
	}
}

func (d *campaignDelivery) GetCampaigns(c *gin.Context) {
	userID, _ := strconv.Atoi(c.Query(("user_id")))

	response := d.campaignUseCase.GetCampaigns(userID)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *campaignDelivery) GetCampaign(c *gin.Context) {
	var campaignUri dto.CampaignUri

	if err := c.ShouldBindUri(&campaignUri); err != nil {
		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"URI validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := d.campaignUseCase.GetCampaign(campaignUri)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *campaignDelivery) CreateCampaign(c *gin.Context) {
	var request dto.CampaignRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if authenticatedUser, ok := c.MustGet("authenticatedUser").(entity.User); ok {
		request.User = authenticatedUser
	} else {
		response := dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			"not authenticated",
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := d.campaignUseCase.CreateCampaign(request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *campaignDelivery) UpdateCampaign(c *gin.Context) {
	var request dto.CampaignRequest
	var CampaignID dto.CampaignUri

	if err := c.ShouldBindUri(&CampaignID); err != nil {
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if authenticatedUser, ok := c.Value("authenticatedUser").(entity.User); ok {
		request.User = authenticatedUser
	} else {
		response := dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			"not authenticated",
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	response := d.campaignUseCase.UpdateCampaign(CampaignID, request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (d *campaignDelivery) CreateCampaignImage(c *gin.Context) {
	var request dto.CampaignImageRequest

	if err := c.ShouldBind(&request); err != nil {
		if errors.Is(err, io.EOF) {
			response := dto.BuildResponse(
				"Body request bind failed",
				"FAILED",
				http.StatusBadRequest,
				map[string]interface{}{"errors": err.Error()},
			)

			c.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		errors := utils.ValidationFormatter(err)
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": errors},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("campaign_image")
	if err != nil {
		response := dto.BuildResponse(
			"Campaign image upload failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authenticatedUser, ok := c.Value("authenticatedUser").(entity.User)
	request.User = authenticatedUser
	if !ok {
		response := dto.BuildResponse(
			"Authentication failed",
			"FAILED",
			http.StatusUnauthorized,
			"not authenticated",
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	fullPath := fmt.Sprintf("images/%d-%s", authenticatedUser.ID, file.Filename)
	if err = c.SaveUploadedFile(file, fullPath); err != nil {
		response := dto.BuildResponse(
			"Avatar upload failed",
			"FAILED",
			http.StatusBadRequest,
			map[string]interface{}{"errors": err.Error()},
		)

		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := d.campaignUseCase.CreateCampaignImage(request, fullPath)
	if response.Meta.Code != http.StatusOK {
		c.AbortWithStatusJSON(response.Meta.Code, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
