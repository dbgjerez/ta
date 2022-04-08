package interfaces

import (
	"net/http"
	"ta-bf-candle-reader/adapter"
	"ta-bf-candle-reader/domain/dto"

	"github.com/gin-gonic/gin"
)

type HealthcheckHandler struct {
}

func HealthcheckGetHandler(wsConnection *adapter.BitfinexConnection) func(c *gin.Context) {
	return func(c *gin.Context) {
		h := dto.Health{}
		if wsConnection.IsConnected() {
			h.Status = dto.HealhStatusUp
			c.JSON(http.StatusOK, h)
		} else {
			h.Status = dto.HealhStatusDown
			c.JSON(http.StatusInternalServerError, h)
		}
	}
}
