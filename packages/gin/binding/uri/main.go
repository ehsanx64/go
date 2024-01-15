package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/:name/:id", func(c *gin.Context) {
		var p Person

		if err := c.ShouldBindUri(&p); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"name": p.Name,
			"uuid": p.ID,
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
