package handlers

import (
	"fmt"
	"net/http"

	jwtauth "github.com/fishmanDK"
	"github.com/gin-gonic/gin"
)

func (h *Handlers) authentication(c *gin.Context) {
	var inp jwtauth.User
	if err := c.BindJSON(&inp); err != nil{

		return
	}
	fmt.Println(inp)
	res, err := h.Service.Authentication.Authentication(inp)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(res, "aaaaa")
	c.JSON(http.StatusOK, res)
}

func (h *Handlers) createUser(c *gin.Context) {

}