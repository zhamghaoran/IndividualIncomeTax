package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	router := gin.Default()
	router.StaticFile("/", "index.html")
	router.POST("/calculate-tax", func(c *gin.Context) {
		// 从请求参数中获取收入金额
		income := c.PostForm("income")

		// 调用计算个人所得税的函数
		tax := calculateTax(income)

		// 返回计算结果
		c.JSON(http.StatusOK, gin.H{
			"tax": tax,
		})
	})

	router.Run(":8080")
}

func calculateTax(income string) float64 {
	money, _ := strconv.ParseFloat(income, 2)
	if money <= 5000 {
		return money
	} else if money <= 8000 && money > 5000 {
		return money * 0.03
	} else if money <= 17000 && money > 8000 {
		return money * 0.10
	} else if money <= 30000 && money > 17000 {
		return money * 0.20
	}
	return money
}
