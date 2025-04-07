package goods

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Good struct {
	ID    int
	Name  string
	State string
}

func (api *API) GetGoods(c *gin.Context) {
	goodsDomain, err := api.service.GetGoods(c)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "error",
			"error":  "there are no goods",
		})
		api.logger.Err(err)
		return
	}
	var goods []Good
	for _, good := range goodsDomain {
		goods = append(goods, Good{
			ID:   good.ID,
			Name: good.Name,
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"goods":  goods,
	})
}
