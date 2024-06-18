package api

import (
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ArticleAPI struct{}

func (b *ArticleAPI) QueryAritcleList(c *gin.Context) {
	articleDto := Req.ArticleDto{}
	err := c.ShouldBindJSON(&articleDto)

	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	list, total, err := articleService.QueryArticleList(articleDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	Res.OkPage(total, list, c)
}

func (b *ArticleAPI) CreateArticle(c *gin.Context) {
	var articleDto Req.ArticleDto
	err := c.ShouldBindJSON(&articleDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	validate := validation.Validation{}
	vr, err := validate.Valid(&articleDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	if !vr {
		Res.ErrorWithMsg("invalid param", c)
		return
	}
	id, err := articleService.CreateArticle(articleDto)
	if err != nil {
		Res.ErrorWithMsg("create article failed", c)
		return
	}
	Res.OkData(id, c)
}

func (b *ArticleAPI) UpdateArticle(c *gin.Context) {
	var articleDto Req.ArticleDto
	err := c.ShouldBindJSON(&articleDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	validate := validation.Validation{}
	vr, _ := validate.Valid(&articleDto)
	if !vr {
		Res.ErrorWithMsg("invalid param", c)
		return
	}
	count, err := articleService.UpdateArticle(articleDto)
	if err != nil {
		Res.ErrorWithMsg("update article failed", c)
		return
	}
	Res.OkData(count, c)
}

func (b *ArticleAPI) DeleteArticle(c *gin.Context) {
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
	count, err := articleService.DelByID(int64(convID))
	if count <= 0 || err != nil {
		Res.ErrorWithMsg("delete article failed", c)
		return
	}
	Res.OkData(count, c)
}
