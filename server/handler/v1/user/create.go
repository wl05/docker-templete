package user

import (
	"fmt"
	. "server/handler/v1"
	"server/model"
	"github.com/gin-gonic/gin"
	"server/pkg/auth"
	"server/pkg/errno"
	"time"
	"github.com/globalsign/mgo/bson"
)

type createResponse struct {
	Username string `json:"username"`
}
type createRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body createRequest true "Create a new user"
// @Success 200 {object} createResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user [post]
func Create(c *gin.Context) {
	var r createRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{
		ID:        bson.NewObjectId(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Username:  r.Username,
		Email:     r.Email,
		Avatar:    "https://user-gold-cdn.xitu.io/2019/5/29/16b028263cf8b532?imageView2/1/w/100/h/100/q/85/format/webp/interlace/1",
		Password:  r.Password,
	}

	// Encrypt the user password.
	encryptedPassword, err := auth.Encrypt(u.Password)
	if err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}
	u.Password = encryptedPassword

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		fmt.Print(err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := createResponse{
		Username: r.Username,
	}

	// Show the user information.
	SendResponse(c, nil, rsp)
}
