package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type PageResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
}

const (
	SUCCESS     = true
	SuccessCode = 10000
	SuccessMsg  = "success"

	ERROR     = false
	ErrorCode = 9999
	ErrorMsg  = "error"
)

func WriteResult(success bool, code int, msg string, data interface{}, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		success,
		code,
		msg,
		data,
	})
}
func Error(c *gin.Context) {
	WriteResult(ERROR, ErrorCode, ErrorMsg, map[string]interface{}{}, c)
}
func ErrorWithMsg(msg string, c *gin.Context) {
	WriteResult(ERROR, ErrorCode, msg, map[string]interface{}{}, c)
}

func ErrorCodeMsg(code int, msg string, c *gin.Context) {
	WriteResult(ERROR, code, msg, map[string]interface{}{}, c)
}
func Ok(c *gin.Context) {
	WriteResult(SUCCESS, SuccessCode, SuccessMsg, map[string]interface{}{}, c)
}
func OkMsg(msg string, c *gin.Context) {
	WriteResult(SUCCESS, SuccessCode, msg, map[string]interface{}{}, c)
}
func OkCodeMsg(code int, msg string, c *gin.Context) {
	WriteResult(SUCCESS, code, msg, map[string]interface{}{}, c)
}
func OkData(data interface{}, c *gin.Context) {
	WriteResult(SUCCESS, SuccessCode, SuccessMsg, data, c)
}
func OkPage(total int64, data interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, PageResponse{
		SUCCESS,
		SuccessCode,
		SuccessMsg,
		data,
		total,
	})
}
func ErrorPage(c *gin.Context) {
	c.JSON(http.StatusOK, PageResponse{
		SUCCESS,
		SuccessCode,
		SuccessMsg,
		map[string]interface{}{},
		0,
	})
}
func OkResponse(response Response, c *gin.Context) {
	c.JSON(http.StatusOK, response)
}
