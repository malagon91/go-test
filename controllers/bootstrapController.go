package controllers

import (
	"api-bootstrap-echo/models"
	"api-bootstrap-echo/repositories"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// NewBootstrapController :
func NewBootstrapController(repo *repositories.BootstrapRepository) *BootstrapController {
	return &BootstrapController{Repository: repo}
}

// BootstrapController :
type BootstrapController struct {
	Repository repositories.IBootstrapRepository
}

// Get
func (controller *BootstrapController) Get(c echo.Context) error {
	id, errConv := strconv.Atoi(c.Param("id"))
	if errConv != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: "ERROR.REQUEST", Message: models.ErrorIDMustBeInteger})
	}

	bootstrap, _ := controller.Repository.Get(id)
	return c.JSON(http.StatusOK, bootstrap)
}

// Post :
func (controller *BootstrapController) Post(c echo.Context) error {
	body, err := BodyToStruct(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ErrorResponse{Code: models.ErrorRequest, Message: models.ErrorDeserializeBody})
	}

	errvalidateModel := ValidateInputModel(body)
	if errvalidateModel.Errors != nil {
		return c.JSON(http.StatusBadRequest, errvalidateModel)
	}

	// TBD: Add repository call.

	return c.JSON(http.StatusCreated, body)
}
