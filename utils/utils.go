package utils

import (
	"math"
	"receipt-processor/models"
	"strconv"
	"strings"
	"time"
)

func CalculatePoints(receipt models.Receipt) int {
    points := 0

    // One point for every alphanumeric character in the retailer name.
    points += len(strings.ReplaceAll(strings.ReplaceAll(receipt.Retailer, " ", ""), "&", ""))

    // 50 points if the total is a round dollar amount with no cents.
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if total == float64(int(total)) {
            points += 50
        }
    }

    // 25 points if the total is a multiple of 0.25.
    if total, err := strconv.ParseFloat(receipt.Total, 64); err == nil {
        if int(total*100)%25 == 0 {
            points += 25
        }
    }

    // 5 points for every two items on the receipt.
    points += (len(receipt.Items) / 2) * 5

    // Points for item description length
    for _, item := range receipt.Items {
        descLen := len(strings.TrimSpace(item.ShortDescription))
        if descLen%3 == 0 {
            if price, err := strconv.ParseFloat(item.Price, 64); err == nil {
                points += int(math.Ceil(price * 0.2))
            }
        }
    }

    // 6 points if the day in the purchase date is odd.
    if date, err := time.Parse("2006-01-02", receipt.PurchaseDate); err == nil {
        if date.Day()%2 != 0 {
            points += 6
        }
    }

    // 10 points if the time of purchase is after 2:00pm and before 4:00pm.
    if t, err := time.Parse("15:04", receipt.PurchaseTime); err == nil {
        if (t.Hour() == 14 && t.Minute() > 0) || (t.Hour() == 15 && t.Minute() < 60) {
            points += 10
        }
    
    }

    return points
}
