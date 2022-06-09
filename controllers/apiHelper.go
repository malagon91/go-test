package controllers

import (
	utils "api-bootstrap-echo/libs"
	"api-bootstrap-echo/models"

	"github.com/labstack/echo/v4"
	"gopkg.in/go-validator/validator.v2"
)

// BodyToStruct :
func BodyToStruct(c echo.Context) (models.Bootstrap, error) {
	bootstrapSample := new(models.Bootstrap)
	err := c.Bind(bootstrapSample)
	return *bootstrapSample, err
}

// ValidateInputModel :
//https://www.thepolyglotdeveloper.com/2019/03/validating-data-structures-variables-golang/
//https://medium.com/@apzuk3/input-validation-in-golang-bc24cdec1835
//https://github.com/go-validator/validator
//https://godoc.org/gopkg.in/go-playground/validator.v2
//https://godoc.org/gopkg.in/go-validator/validator.v2
func ValidateInputModel(bootstrapSample models.Bootstrap) (errToOut models.ErrorsResponse) {
	errValidator := validator.NewValidator().Validate(bootstrapSample)
	if errValidator != nil {
		errToOut.Errors = utils.AppendErrors(errValidator)
	}
	return errToOut
}
