package utils

import (
	"fmt"
	"math"
	"time"
)

func CalculateDuration(d string) string {
	convertedTime, _ := time.Parse("20060102T150405.000Z", d)
	now := time.Now()
	timeSinceLastPlayed := now.Sub(convertedTime)
	return formatDuration(timeSinceLastPlayed) + " ago"
}

// formatDuration converts a time.Duration to a human-readable string
func formatDuration(d time.Duration) string {
	days := d / (24 * time.Hour)
	d -= days * 24 * time.Hour
	hours := d / time.Hour
	return fmt.Sprintf("%d days and %d hours", days, hours)
}

func RoundToNearestTwoDecimals(num float64) float64 {
	return math.Round(num*100) / 100
}
