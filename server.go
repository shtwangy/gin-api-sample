package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name  string `form:"name"`
	Email string `form:"email"`
}

func main() {
	r := gin.Default()
	r.GET("/user", show)
	r.POST("/users", display)

	r.Run(":1323")
}

func show(c *gin.Context) {
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, u)
}

func display(c *gin.Context) {
	u := User{}
	err := c.Bind(&u)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusOK, u)
}
