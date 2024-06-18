package api

import (
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TagAPI struct{}

func (base *BaseAPI) QueryTagList(c *gin.Context) {
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

func (base *BaseAPI) CreateTag(context *gin.Context) {
	var tagDto Req.TagDto
	err := context.ShouldBindJSON(&tagDto)
	if err != nil {
		Res.ErrorWithMsg("invalid param", context)
	}

	id, err := tagService.CreateTag(tagDto)
	if err != nil {
		Res.ErrorWithMsg("create tag failed", context)
		return
	}
	Res.OkData(id, context)
}

func (base *BaseAPI) DeleteTag(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		Res.ErrorWithMsg("id can't be empty", c)
		return
	}
	convID, err := strconv.Atoi(id)
	if err != nil {
		Res.ErrorWithMsg("invalid id ", c)
		return
	}
	count, err := categoryService.DelByID(int64(convID))
	if count <= 0 || err != nil {
		Res.ErrorWithMsg("delete tag failed", c)
		return
	}
	Res.OkData(count, c)
}
