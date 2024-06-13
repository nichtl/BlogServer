package api

import (
	Req "blogServe/business/model/request"
	Res "blogServe/business/model/response"
	"github.com/gin-gonic/gin"
)

type ArticleApi struct{}

func (b *ArticleApi) QueryAritcleList(c *gin.Context) {
	articleDto := Req.ArticleDto{}
	err := c.ShouldBindJSON(&articleDto)

	if err != nil {
		Res.OkData(articleDto, c)
		return
	}
	list, total, err := articleService.QueryAritcleList(articleDto)
	if err != nil {
		Res.ErrorWithMsg(err.Error(), c)
		return
	}
	Res.OkPage(total, list, c)
}
