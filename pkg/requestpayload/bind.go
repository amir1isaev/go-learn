package requestpayload

import (
	"encoding/json"
	"go-learn/pkg/validator"

	"github.com/gin-gonic/gin"
)

func Bind(c *gin.Context, obj any) error {
	err := json.NewDecoder(c.Request.Body).Decode(&obj)
	if err != nil {
		return err
	}

	validate := validator.Get()
	err = validate.StructCtx(c, obj)
	if err != nil {
		return err
	}

	return nil
}
