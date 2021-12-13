package helper

import "github.com/go-playground/validator/v10"

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(
	message string,
	code int,
	status string,
	data interface{},
) Response {
	metaData := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}

	responseData := Response{
		Meta: metaData,
		Data: data,
	}

	return responseData
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
