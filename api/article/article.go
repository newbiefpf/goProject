package article

import (
	"github.com/gin-gonic/gin"
	"goProject/db"
	"goProject/model/dbModel"
	"goProject/utility/ginResult"
	"strconv"
)

func GetArticleList(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))
	offsetIndex := (page - 1) * pageSize
	var article []dbModel.Article
	var DB = db.ConnectDb
	var Count int32
	//查数据总数Offset(-1).Limit(-1).Count(&Count)
	result := DB.Limit(pageSize).Offset(offsetIndex).Find(&article).Offset(-1).Limit(-1).Count(&Count)

	if result.Error != nil {
		c.JSON(200, ginResult.OK.WithMsg("服务器错误!"))
	} else {
		res := struct {
			TotalCount int32 `json:"totalCount"`
			Data       interface{}
		}{
			TotalCount: Count,
			Data:       article,
		}
		c.JSON(200, ginResult.OK.WithData(res))
	}

}

func CreatArticle(c *gin.Context) {
	var article dbModel.Article
	var DB = db.ConnectDb
	_ = c.ShouldBindJSON(&article)
	if article.Title == "" {
		c.JSON(200, ginResult.ErrParam.WithMsg("标题不能为空"))
		return
	} else if article.ContentHtml == "" {
		c.JSON(200, ginResult.ErrParam.WithMsg("类容不能为空"))
		return
	} else if article.UserID == 0 {
		c.JSON(200, ginResult.ErrParam.WithMsg("UserID不能为空"))
		return
	}
	result := DB.Create(&article)
	if result.Error != nil {
		c.JSON(200, ginResult.OK.WithMsg("服务器错误!"))
	} else {
		c.JSON(200, ginResult.OK.WithMsg("发表成功等待管理员审核!"))
	}

}
