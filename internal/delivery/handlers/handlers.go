package handlers

import (
	"go-learn/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	userUC usecase.User
}

func NewHandlers(userUC usecase.User) *Handlers {
	return &Handlers{
		userUC: userUC,
	}
}

func (h Handlers) MapHandlers(engine *gin.Engine) {
	apiGroup := engine.Group("/api")

	userGroup := apiGroup.Group("/user")

	h.userRoutes(userGroup)

}

func (h Handlers) userRoutes(router *gin.RouterGroup) {
	router.POST("/", h.CreateUser)
}
