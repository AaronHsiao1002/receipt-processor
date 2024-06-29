package utils

import (
	"receipt-processor/models"
	"testing"
)

func TestCalculatePoints(t *testing.T) {
    tests := []struct {
        name     string
        receipt  models.Receipt
        expected int
    }{
        {
            name: "Provided example 1",
            receipt: models.Receipt{
                Retailer:     "Target",
                PurchaseDate: "2022-01-01",
                PurchaseTime: "13:01",
                Items: []models.Item{
                    {ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
                    {ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
                    {ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
                    {ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
                    {ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
                },
                Total: "35.35",
            },
            expected: 28,
        },
        {
            name: "Provided example 2",
            receipt: models.Receipt{
                Retailer:     "M&M Corner Market",
                PurchaseDate: "2022-03-20",
                PurchaseTime: "14:33",
                Items: []models.Item{
                    {ShortDescription: "Gatorade", Price: "2.25"},
                    {ShortDescription: "Gatorade", Price: "2.25"},
                    {ShortDescription: "Gatorade", Price: "2.25"},
                    {ShortDescription: "Gatorade", Price: "2.25"},
                },
                Total: "9.00",
            },
            expected: 109,
        },
		{
            name: "Round dollar total",
            receipt: models.Receipt{
                Retailer:     "Walmart",
                PurchaseDate: "2022-01-02",
                PurchaseTime: "12:30",
                Items: []models.Item{
                    {ShortDescription: "Item A", Price: "5.00"},
                    {ShortDescription: "Item B", Price: "10.00"},
                },
                Total: "15.00",
            },
            expected: 90, // 7 + 50 + 25 + 5 + 3
        },
		
        {
            name: "Total multiple of 0.25",
            receipt: models.Receipt{
                Retailer:     "CVS",
                PurchaseDate: "2022-01-02",
                PurchaseTime: "09:00",
                Items: []models.Item{
                    {ShortDescription: "Item C", Price: "2.25"},
                    {ShortDescription: "Item D", Price: "4.50"},
                },
                Total: "6.75",
            },
            expected: 35, // 3 (retailer) + 25 (multiple of 0.25) + 5 (2 items) + 2 (description length multiple of 3, rounded up)
        },
        {
            name: "Purchase date with odd day",
            receipt: models.Receipt{
                Retailer:     "Best Buy",
                PurchaseDate: "2022-01-03",
                PurchaseTime: "10:00",
                Items: []models.Item{
                    {ShortDescription: "Item E", Price: "7.99"},
                    {ShortDescription: "Item F", Price: "14.99"},
                },
                Total: "22.98",
            },
            expected: 23, // 7 (retailer) + 6 (odd day) + 5 (2 items) + 5
        },
        {
            name: "Purchase time between 2:00pm and 4:00pm",
            receipt: models.Receipt{
                Retailer:     "Amazon",
                PurchaseDate: "2022-01-02",
                PurchaseTime: "15:32",
                Items: []models.Item{
                    {ShortDescription: "ItemG", Price: "9.99"},
                },
                Total: "9.99",
            },
            expected: 16, // 6 (retailer) + 10 (time between 2pm and 4pm)
        },
        {
            name: "Item description length multiple of 3",
            receipt: models.Receipt{
                Retailer:     "eBay",
                PurchaseDate: "2022-01-02",
                PurchaseTime: "14:00",
                Items: []models.Item{
                    {ShortDescription: "123", Price: "3.00"},
                },
                Total: "3.00",
            },
            expected: 80, 
        },
		{
			// 5 points - retailer name
			// 10 points - purchased between 2pm and 4pm
			// total 15 points
			name: "Purchase time between 14:01 and 15:59",
			receipt: models.Receipt{
				Retailer:     "Store",
				PurchaseDate: "2022-01-02",
				PurchaseTime: "14:30",
				Items: []models.Item{
					{ShortDescription: "Item1", Price: "1.99"},
				},
				Total: "1.99",
			},
			expected: 15,
		},

    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := CalculatePoints(tt.receipt); got != tt.expected {
                t.Errorf("CalculatePoints() = %v, want %v", got, tt.expected)
            }
        })
    }
}
