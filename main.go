package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "time"
)

var receipts = make(map[string]Receipt)
var points = make(map[string]int)

func main() {
    r := gin.Default()
    r.POST("/receipts/process", ProcessReceipt)
    r.GET("/receipts/:id/points", GetPoints)

    r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func ProcessReceipt(c *gin.Context) {
    var receipt Receipt
    if err := c.BindJSON(&receipt); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    id := generateID()
    receipts[id] = receipt
    points[id] = calculatePoints(receipt)

    c.JSON(http.StatusOK, gin.H{"id": id})
}

func GetPoints(c *gin.Context) {
    id := c.Param("id")
    if _, exists := points[id]; !exists {
        c.JSON(http.StatusNotFound, gin.H{"error": "ID not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"points": points[id]})
}
