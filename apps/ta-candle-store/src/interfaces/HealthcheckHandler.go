package interfaces

import (
	"net/http"
	"ta-candle-store/domain/dto"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
}

func HealthcheckGetHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Health{}
		h.Status = dto.HealhStatusUp
		c.JSON(http.StatusOK, h)
	}
}
