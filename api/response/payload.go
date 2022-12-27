package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	respOk   = 0  // "OK"
	respFail = -1 // "FAIL"
)

type RespResult struct {
	Data interface{} `json:"data"`
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
}

func Success(c *gin.Context, resp *RespResult) {
	resp.Code = respOk
	resp.Msg = "success"

	c.JSON(http.StatusOK, resp)
}

// Fail . return 400 and error msg
func Fail(c *gin.Context, e error) {
	c.JSON(http.StatusBadRequest, RespResult{
		Code: respFail,
		Msg:  e.Error(),
	})
	c.Abort()
}
