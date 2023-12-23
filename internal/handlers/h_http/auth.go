package h_http

import (
	"fmt"
	"net/http"

	jwtauth "github.com/fishmanDK"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) authentication(c *gin.Context) {
	var inp jwtauth.User
	if err := c.BindJSON(&inp); err != nil {

		return
	}
	res, err := h.Service.Authentication.Authentication(inp)
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(http.StatusOK, res)
}

func (h *Handlers) createUser(c *gin.Context) {
	var inp jwtauth.CreateUser
	if err := c.BindJSON(&inp); err != nil {

		return
	}
	user_id, err := h.Service.Authentication.CreateUser(inp)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(user_id, "aaa")
	type responce struct {
		User_id int64
	}
	c.JSON(http.StatusOK, &responce{
		User_id: user_id,
	})
}
