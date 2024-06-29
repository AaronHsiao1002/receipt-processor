package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    r.POST("/receipts/process", processReceiptHandler)
    r.GET("/receipts/:id/points", getPointsHandler)

    r.Run(":8080")
}
