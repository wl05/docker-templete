package user

import (
	. "server/handler/v1"
	"server/model"
	"server/pkg/auth"
	"server/pkg/errno"
	"server/pkg/token"
	"server/config"
	"fmt"
	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginResponse struct {
	Username string `json:"username"`
}

type tokenResponse struct {
	Token string `json:"token"`
}


// @Summary Login generates the authentication token
// @Description 登录
// @Tags login
// @Accept  json
// @Produce  json
// @Param user body user.loginRequest true "login"
// @Success 200 {object} user.loginResponse "{"code":0,"message":"OK","data":{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE1MjgwMTY5MjIsImlkIjowLCJuYmYiOjE1MjgwMTY5MjIsInVzZXJuYW1lIjoiYWRtaW4ifQ.LjxrK9DuAwAzUD8-9v43NzWBN7HXsSLfebw92DKd1JQ"}}"
// @Router /v1/login [post]
func Login(c *gin.Context) {
	// Binding the data with the user struct.
	var u loginRequest
	if err := c.Bind(&u); err != nil {
		fmt.Println("err", err)
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Get the user information by the login username.
	d, err := model.GetUserByName(u.Username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	// Compare the login password with the user password.
	if err := auth.Compare(d.Password, u.Password); err != nil {
		SendResponse(c, errno.ErrPasswordIncorrect, nil)
		return
	}

	// Sign the json web token.
	t, err := token.Sign(c, token.Context{ID: d.ID.Hex(), Username: d.Username}, "")
	if err != nil {
		SendResponse(c, errno.ErrToken, nil)
		return
	}
	// 种下cookie
	c.SetCookie("xtoken", t, config.CooKieMaxAge * 60 * 60,"/", "", false, true)
	SendResponse(c, nil, tokenResponse{Token: t})
}
