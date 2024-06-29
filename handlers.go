package main

import (
	"net/http"
	"receipt-processor/models"
	"receipt-processor/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var receipts = make(map[string]models.Receipt)

func processReceiptHandler(c *gin.Context) {
    var receipt models.Receipt

    if err := c.ShouldBindJSON(&receipt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    receipt.ID = uuid.New().String()
    receipt.Points = utils.CalculatePoints(receipt)
    receipts[receipt.ID] = receipt

    c.JSON(http.StatusOK, gin.H{"id": receipt.ID})
}

func getPointsHandler(c *gin.Context) {
    id := c.Param("id")

    receipt, exists := receipts[id]
    if !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "receipt not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"points": receipt.Points})
}
