package handlers

import (
	"go-learn/internal/delivery/http/dto"
	response "go-learn/internal/delivery/http/response"
	"go-learn/pkg/requestpayload"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handlers) CreateUser(gCtx *gin.Context) {
	var (
		reqBody dto.CreateUserDTO
	)

	if err := requestpayload.Bind(gCtx, &reqBody); err != nil {
		_ = gCtx.Error(err)
		return
	}

	user, err := h.userUC.Create(gCtx, reqBody)

	if err != nil {
		_ = gCtx.Error(err)
		return
	}

	gCtx.JSON(http.StatusOK, response.NewResponse(user))
}
