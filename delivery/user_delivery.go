package delivery

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"service-campaign-startup/model/dto"
	"service-campaign-startup/model/entity"
	"service-campaign-startup/usecase"
	"service-campaign-startup/utils"

	"github.com/gin-gonic/gin"
)

type userDelivery struct {
	userUseCase usecase.UserUseCase
}

func NewUserDelivery(userUseCase usecase.UserUseCase) UserDelivery {
	return &userDelivery{
		userUseCase: userUseCase,
	}
}

func (deliveries *userDelivery) RegisterUser(c *gin.Context) {
	var request dto.UserRegisterRequest

	errBind := c.ShouldBindJSON(&request)
	if errBind != nil && errors.Is(errBind, io.EOF) {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if errBind != nil {
		errors := utils.ValidationFormatter(errBind)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := deliveries.userUseCase.RegisterUser(request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (deliveries *userDelivery) LoginUser(c *gin.Context) {
	var request dto.UserLoginRequest

	errBind := c.ShouldBindJSON(&request)
	if errBind != nil && errors.Is(errBind, io.EOF) {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if errBind != nil {
		errors := utils.ValidationFormatter(errBind)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := deliveries.userUseCase.LoginUser(request)
	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if response.Meta.Code == http.StatusUnauthorized {
		c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (deliveries *userDelivery) GetUserByEmail(c *gin.Context) {
	var request dto.EmailCheckRequest

	errBind := c.ShouldBindJSON(&request)
	if errBind != nil && errors.Is(errBind, io.EOF) {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if errBind != nil {
		errors := utils.ValidationFormatter(errBind)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isEmailExist, err := deliveries.userUseCase.GetUserByEmail(request)
	if err != nil {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if isEmailExist {
		response := dto.BuildResponse(
			"Email registration faled",
			"FAILED",
			http.StatusUnprocessableEntity,
			"Email already registered",
		)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := dto.BuildResponse(
		"Email has retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		"Email is available",
	)
	c.JSON(http.StatusOK, response)
}

func (deliveries *userDelivery) GetUserById(c *gin.Context) {
	var request dto.EmailCheckRequest

	errBind := c.ShouldBindJSON(&request)
	if errBind != nil && errors.Is(errBind, io.EOF) {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Body request bind failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	if errBind != nil {
		errors := utils.ValidationFormatter(errBind)
		err := map[string]interface{}{"ERROR": errors}
		response := dto.BuildResponse(
			"Body request validation failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	isEmailExist, err := deliveries.userUseCase.GetUserByEmail(request)
	if err != nil {
		err := map[string]interface{}{"ERROR": errBind.Error()}
		response := dto.BuildResponse(
			"Database query error or database connection problem",
			"FAILED",
			http.StatusInternalServerError,
			err,
		)
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if isEmailExist {
		response := dto.BuildResponse(
			"Email registration faled",
			"FAILED",
			http.StatusUnprocessableEntity,
			"Email already registered",
		)
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
		return
	}

	response := dto.BuildResponse(
		"Email has retrieved successfully",
		"SUCCESS",
		http.StatusOK,
		"Email is available",
	)
	c.JSON(http.StatusOK, response)
}

func (deliveries *userDelivery) SaveUserAvatar(c *gin.Context) {
	file, err := c.FormFile("AVATAR")
	if err != nil {
		log.Println(err.Error())
		err := map[string]interface{}{"ERROR": err.Error()}
		response := dto.BuildResponse(
			"Avatar upload failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authenticatedUser := c.MustGet("authenticatedUser").(entity.User)
	fullPath := fmt.Sprintf("images/%d-%s", authenticatedUser.ID, file.Filename)
	err = c.SaveUploadedFile(file, fullPath)
	if err != nil {
		log.Println(err.Error())
		err := map[string]interface{}{"ERROR": err.Error()}
		response := dto.BuildResponse(
			"Avatar upload failed",
			"FAILED",
			http.StatusBadRequest,
			err,
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	response := deliveries.userUseCase.SaveUserAvatar(1, fullPath)
	if response.Meta.Code == http.StatusNotFound {
		c.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	if response.Meta.Code == http.StatusInternalServerError {
		c.AbortWithStatusJSON(http.StatusInternalServerError, response)
		return
	}

	if response.Meta.Code == http.StatusBadRequest {
		c.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	c.JSON(http.StatusOK, response)
}
