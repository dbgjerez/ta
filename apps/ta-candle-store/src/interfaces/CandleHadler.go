package interfaces

import (
	"net/http"
	"ta-candle-store/domain/model"

	"github.com/gin-gonic/gin"
)

const (
	ParamIDName  = "id" // find by id
	BodyDataName = "data"
)

type CandleController struct {
	repository *model.CandleRepository
}

func NewCandleController(dao *model.CandleRepository) *CandleController {
	return &CandleController{repository: dao}
}

func (controller *CandleController) GetCandle(c *gin.Context) {
	id := c.Param(ParamIDName)
	candles := controller.repository.FindAllByType(id)
	c.JSON(http.StatusOK, gin.H{BodyDataName: candles})
}
