package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrResponse struct {
	Err     int    `json:"err"`
	Message string `json:"message"`
}

type SuccessResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}

func HandleErrResponse(c *gin.Context, err error) {
	resp := ErrResponse{
		Err:     500,
		Message: err.Error(),
	}
	c.JSON(http.StatusInternalServerError, resp)
}

func HandleSuccessResponse[T any](c *gin.Context, data T) {
	resp := SuccessResponse[T]{
		Code:    0,
		Message: "susscess",
		Data:    data,
	}
	c.JSON(http.StatusOK, resp)
}
