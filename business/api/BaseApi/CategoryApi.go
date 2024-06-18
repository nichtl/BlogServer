package api

import (
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryAPI struct{}

func (base *BaseAPI) QueryAllFirstCategory(context *gin.Context) {
	categorys, err := categoryService.FindFirstAll()
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(categorys, context)
	return
}

func (base *BaseAPI) QueryChildrenByID(context *gin.Context) {
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
	category, err := categoryService.FindChildrenByFatherID(v)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(category, context)
}

func (base *BaseAPI) QueryCategoryByID(context *gin.Context) {
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
	category, err := categoryService.FindCategoryByID(v)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), context)
		return
	}
	Res.OkData(category, context)
	return
}

func (base *BaseAPI) CreateCategory(context *gin.Context) {
	var categoryDto Req.CategoryDto
	err := context.ShouldBindJSON(&categoryDto)
	if err != nil {
		Res.ErrorWithMsg("invalid param", context)
	}

	id, err := categoryService.CreateCategory(categoryDto)
	if err != nil {
		Res.ErrorWithMsg("create article failed", context)
		return
	}
	Res.OkData(id, context)
}

func (base *BaseAPI) DeleteCategory(c *gin.Context) {
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
		Res.ErrorWithMsg("delete category failed", c)
		return
	}
	Res.OkData(count, c)
}
