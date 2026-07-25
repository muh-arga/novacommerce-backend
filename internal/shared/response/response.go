package response

import "github.com/gin-gonic/gin"

type SuccessResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
}

func Success(c *gin.Context, status int, data any) {
	c.JSON(status, SuccessResponse{
		Success: true,
		Message: "Success",
		Data:    data,
	})
}

func Error(c *gin.Context, status int, message string, errors any) {
	c.JSON(status, ErrorResponse{
		Success: false,
		Message: message,
		Errors:  errors,
	})
}
