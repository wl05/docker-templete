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

// @Summary signup
// @Description signup
// @Tags signup
// @Accept  json
// @Produce  json
// @Param user body createRequest true "signup"
// @Success 200 {object} createResponse "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/signup [post]
func Signup(c *gin.Context) {
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
		Avatar:    "http://m.imeitou.com/uploads/allimg/2019022710/b4jgfd5plcg.jpg",
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
