package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"exchangeapp/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) //绑定参数
		return
	}
	hashedPwd, error := utils.HashPassword(user.Password)
	if error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": error.Error()})
		return
	}
	user.Password = hashedPwd

	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := global.Db.AutoMigrate(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return //自动创建不存在的表，防止找不到对应的表导致数据库崩溃
	}
	if err := global.Db.Create(&user).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return //Create()：GORM 框架提供的新增数据方法，作用是将传入的结构体数据插入到对应的数据库表中
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token}) //将生成的jwt令牌返回给前端
}

func Login(ctx *gin.Context) {
	var input struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.ShouldBind(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var user models.User
	if err := global.Db.Where("username = ?", input.Username).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
	if !utils.CheckPassword(input.Password, user.Password) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "wrong credentials"})
		return
	}
	token, err := utils.GenerateJWT(user.Username)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
