package main

import (
    "math"
    "regexp"
    "strconv"
    "strings"
    "time"
)

func calculatePoints(receipt Receipt) int {
    points := 0

    points += len(alphanumeric(receipt.Retailer))
    if isRoundDollar(receipt.Total) {
        points += 50
    }
    if isMultipleOfQuarter(receipt.Total) {
        points += 25
    }
    points += (len(receipt.Items) / 2) * 5
    for _, item := range receipt.Items {
        if len(strings.TrimSpace(item.ShortDescription))%3 == 0 {
            price, _ := strconv.ParseFloat(item.Price, 64)
            points += int(math.Ceil(price * 0.2))
        }
    }
    if isOddDay(receipt.PurchaseDate) {
        points += 6
    }
    if isBetweenTwoAndFour(receipt.PurchaseTime) {
        points += 10
    }

    return points
}

func alphanumeric(s string) string {
    re := regexp.MustCompile("[^a-zA-Z0-9]+")
    return re.ReplaceAllString(s, "")
}

func isRoundDollar(total string) bool {
    totalFloat, _ := strconv.ParseFloat(total, 64)
    return totalFloat == float64(int(totalFloat))
}

func isMultipleOfQuarter(total string) bool {
    totalFloat, _ := strconv.ParseFloat(total, 64)
    return math.Mod(totalFloat, 0.25) == 0
}

func isOddDay(date string) bool {
    t, _ := time.Parse("2006-01-02", date)
    return t.Day()%2 != 0
}

func isBetweenTwoAndFour(timeStr string) bool {
    t, _ := time.Parse("15:04", timeStr)
    return t.Hour() == 14 && t.Minute() > 0 || t.Hour() == 15
}
