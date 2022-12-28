package handler

import (
	"fmt"
	"github.com/doge-verse/zk-doge-backend/api/middleware"
	"github.com/doge-verse/zk-doge-backend/api/request"
	"github.com/doge-verse/zk-doge-backend/api/response"
	"github.com/doge-verse/zk-doge-backend/internal/dao"
	"github.com/doge-verse/zk-doge-backend/internal/service/user"
	"github.com/doge-verse/zk-doge-backend/models"
	"github.com/doge-verse/zk-doge-backend/util"

	"github.com/gin-gonic/gin"
)

// Login .
// @Tags auth
// @Summary User login
// @accept application/json
// @Produce application/json
// @Param data body request.Login true "login param"
// @Success 200 {object} response.RespResult{data=response.UserInfo}
// @Router /login [post]
func Login(c *gin.Context) {
	var param request.Login
	if err := c.ShouldBind(&param); err != nil {
		response.Fail(c, fmt.Errorf("param error"))
		return
	}
	if !util.CheckAddr(param.Address, param.Signature, param.SignData) {
		response.Fail(c, fmt.Errorf("signature fail"))
		return
	}
	userInfo, err := user.Srv.GetUserByQuery(dao.Query{
		Address: param.Address,
	})
	if err != nil {
		newUser := &models.User{
			Address: param.Address,
		}
		userInfo, err = user.Srv.UserRegister(newUser)
		if err != nil {
			response.Fail(c, err)
			return
		}
	}
	tokenStr, expiresAt, err := middleware.SignJwt(userInfo.ID)
	if err != nil {
		response.Fail(c, err)
		return
	}
	response.Success(c, &response.RespResult{
		Data: response.UserInfo{
			ID:        userInfo.ID,
			Name:      userInfo.Name,
			Address:   userInfo.Address,
			Email:     userInfo.Email,
			Token:     tokenStr,
			ExpiresAt: expiresAt,
		},
	})
}

func getUserID(c *gin.Context) uint {
	claims, _ := c.Get("claims")
	return claims.(*request.CustomClaims).UserID
}
