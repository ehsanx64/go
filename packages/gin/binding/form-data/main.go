package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a" binding:"required"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b" binding:"required"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

func main() {
	r := gin.Default()

	r.GET("/getb", getB)
	r.GET("/getc", getC)
	r.GET("/getd", getD)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func getB(c *gin.Context) {
	var b StructB

	err := c.Bind(&b)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func getC(c *gin.Context) {
	var sc StructC

	c.Bind(&sc)
	c.JSON(http.StatusOK, gin.H{
		"a": sc.NestedStructPointer,
		"c": sc.FieldC,
	})
}

func getD(c *gin.Context) {
	var d StructD

	c.Bind(&d)
	c.JSON(http.StatusOK, gin.H{
		"x": d.NestedAnonStruct,
		"d": d.FieldD,
	})
}
