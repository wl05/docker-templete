package user

import (
	. "server/handler/v1"
	"server/model"
	"github.com/gin-gonic/gin"
	"server/pkg/auth"

)
type CreateResponse struct {
	Username string `json:"username"`
}

func Create(c *gin.Context) {

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		Username: r.Username,
		Email: r.Email,
		Avatar:   "https://user-gold-cdn.xitu.io/2019/5/29/16b028263cf8b532?imageView2/1/w/100/h/100/q/85/format/webp/interlace/1",
		Password: r.Password,
	}

	// Encrypt the user password.
	if err := auth.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
