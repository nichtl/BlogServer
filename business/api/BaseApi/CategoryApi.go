package api

import (
	Res "blogServe/business/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryApi struct{}

func (c *BaseApi) QueryAllFirstCategory(context *gin.Context) {
	categorys, err := categoryService.FindFirstAll()
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(categorys, context)
	return
}

func (c *BaseApi) QueryChildrenById(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		Res.ErrorWithMsg("id is required", context)
		return
	}
	v, err := strconv.Atoi(id)
	if err != nil {
		Res.ErrorWithMsg("id is invalid", context)
		return
	}
	category, err := categoryService.FindChildrenByFatherId(v)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(category, context)
	return
}

func (c *BaseApi) QueryCategoryById(context *gin.Context) {
	id := context.Param("id")
	if id == "" {
		Res.ErrorWithMsg("id is required", context)
		return
	}
	v, err := strconv.Atoi(id)
	if err != nil {
		Res.ErrorWithMsg("id is invalid", context)
		return
	}
	category, err := categoryService.FindCategoryById(v)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(category, context)
	return
}
