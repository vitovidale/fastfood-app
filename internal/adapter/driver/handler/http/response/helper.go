package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vitovidale/internal/core/domain"
)

type DefaultResponse struct {
	Data any `json:"data,omitempty"`
}

type ErrorResponse struct {
	Messages []string `json:"messages" example:"Error message 1, Error message 2"`
}

func newResponse(data any) DefaultResponse {
	return DefaultResponse{
		Data: data,
	}
}

func newErrorResponse(messages []string) ErrorResponse {
	return ErrorResponse{
		Messages: messages,
	}
}

func parseError(err error) []string {
	var errMsgs []string

	if errors.As(err, &validator.ValidationErrors{}) {
		for _, err := range err.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.Error())
		}
	} else {
		errMsgs = append(errMsgs, err.Error())
	}

	return errMsgs
}

func HandleSuccess(ctx *gin.Context, data any) {
	rsp := newResponse(data)
	ctx.JSON(http.StatusOK, rsp)
}

var requestErrorStatusMap = map[error]int{
	domain.ErrorInternal:        http.StatusInternalServerError,
	domain.ErrorDataNotFound:    http.StatusNotFound,
	domain.ErrorConflictingData: http.StatusConflict,
}

func HandleBadRequest(ctx *gin.Context, err error) {
	errMsgs := parseError(err)
	errRsp := newErrorResponse(errMsgs)
	ctx.JSON(http.StatusBadRequest, errRsp)
}

func HandleError(ctx *gin.Context, err error) {
	statusCode, ok := requestErrorStatusMap[err]
	if !ok {
		statusCode = http.StatusInternalServerError
	}

	errMsg := parseError(err)
	errRsp := newErrorResponse(errMsg)
	ctx.JSON(statusCode, errRsp)
}
