package controllers

import (
	"exchangeapp/global"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id") //获取动态参数id

	likeKey := "article:" + articleID + ":likes" /*Redis 键名（Key）设计规范
	[实体名] : [唯一ID] : [属性名]*/
	if err := global.RedisDB.Incr(likeKey).Err(); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked article"})

}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Get(likeKey).Result()

	if err == redis.Nil {
		likes = "0"
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"likes": likes})

}
