package controllers

import (
	"exchangeapp/global"
	"exchangeapp/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 完成汇率兑换的功能
func CreateExchangeRate(ctx *gin.Context) {
	var exchangeRate models.ExchangeRate
	if err := ctx.ShouldBindJSON(&exchangeRate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	} //将客户端传来的 HTTP 请求体中的 JSON 数据，自动解析并赋值到 exchangeRate 结构体变量中（模型绑定）
	exchangeRate.Date = time.Now() //把创建这个汇率的时间赋值给date

	if err := global.Db.AutoMigrate(&exchangeRate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return //建表
	}
	if err := global.Db.Create(&exchangeRate).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return //创建记录
	}

	ctx.JSON(http.StatusCreated, exchangeRate)
}

func GetExchangeRate(ctx *gin.Context) {
	var exchangeRates []models.ExchangeRate
	if err := global.Db.Find(&exchangeRates).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, exchangeRates)
}
