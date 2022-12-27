package handler

import (
	"github.com/doge-verse/zk-doge-backend/api/request"
	"github.com/doge-verse/zk-doge-backend/api/response"
	"github.com/doge-verse/zk-doge-backend/internal/service/user"
	"github.com/doge-verse/zk-doge-backend/models"
	"github.com/gin-gonic/gin"
)

// UpdateEmail .
// @Tags auth
// @Summary update user email
// @accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param Authorization header string false "token"
// @Param data body request.UpdateEmail true "update user email"
// @Success 200 {object} response.RespResult
// @Router /user/email [post]
func UpdateEmail(c *gin.Context) {
	param := request.UpdateEmail{}
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Fail(c, err)
		return
	}
	param.UserID = getUserID(c)
	if err := user.Srv.UpdateEmail(param.UserID, param.Email); err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: nil,
	})
}

// Register .
func Register(c *gin.Context) {
	param := models.User{}
	if err := c.ShouldBindQuery(&param); err != nil {
		response.Fail(c, err)
		return
	}
	userInfo, err := user.Srv.UserRegister(&param)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: userInfo,
	})
}
