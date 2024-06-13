package api

import (
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"github.com/gin-gonic/gin"
)

type TagApi struct{}

func (b *BaseApi) QueryTagList(c *gin.Context) {
	tag := Req.TagDto{}
	err := c.ShouldBindJSON(&tag)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}

	list, total, err := tagService.QueryTagList(tag)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	Res.OkPage(total, list, c)
}
