package user

import (
	. "server/handler/v1"
	"server/model"
	"server/pkg/errno"
	"server/pkg/token"
	"github.com/gin-gonic/gin"
)

// @Summary Get an user
// @Description Get an user
// @Tags user
// @Accept  json
// @Produce  json
// @Success 200 {object} model.UserInfo "{"code":0,"message":"OK","data":{"username":"kong"}}"
// @Router /v1/user/info [get]
func GetUserInfo(c *gin.Context) {
	res, _ := token.ParseRequest(c)
	user, err := model.GetUserById(res.ID)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}
	SendResponse(c, nil, user)
}
