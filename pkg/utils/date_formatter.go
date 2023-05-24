// utils/time_formatter.go

package utils

import (
	"fmt"
	"time"
)

const (
	timeLayout = "2006-01-02 15:04:05.999999999 -0700 MST"
)

func parseTime(inputTime string) (time.Time, error) {
	return time.Parse(timeLayout, inputTime)
}

func formatTimeDifference(duration time.Duration, singularUnit string, pluralUnit string) string {
	timeDiff := int(duration.Hours() / 24)
	if timeDiff < 30 {
		if timeDiff == 1 {
			return fmt.Sprintf("%d %s", timeDiff, singularUnit)
		}
		return fmt.Sprintf("%d %s", timeDiff, pluralUnit)
	}

	months := timeDiff / 30
	if months == 1 {
		return fmt.Sprintf("%d month", months)
	}
	return fmt.Sprintf("%d months", months)
}

func FormatExpiresAt(expiresAt string) string {
	expirationTime, err := parseTime(expiresAt)
	if err != nil {
		return "Invalid date"
	}

	remainingDuration := expirationTime.Sub(time.Now())
	if remainingDuration < 0 {
		return "Expired"
	}

	return formatTimeDifference(remainingDuration, "day", "days") + " left"
}

func FormatUsedAt(usedAt string) string {
	lastUsedTime, err := parseTime(usedAt)
	if err != nil {
		return "Invalid date"
	}

	elapsedDuration := time.Since(lastUsedTime)

	if elapsedDuration < time.Minute {
		return "Recently used"
	}

	if elapsedDuration < time.Hour {
		return formatTimeDifference(elapsedDuration, "minute", "minutes") + " ago"
	} else if elapsedDuration < time.Hour*24 {
		return formatTimeDifference(elapsedDuration, "hour", "hours") + " ago"
	}

	return formatTimeDifference(elapsedDuration, "day", "days") + " ago"
}

func FormatCreatedAt(createdAt string) string {
	timestamp, err := parseTime(createdAt)
	if err != nil {
		return "Invalid date"
	}

	return timestamp.Format("2006/01/02")
}
