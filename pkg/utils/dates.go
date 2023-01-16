package utils

import (
	"fmt"
	"time"
)

func GetDate(day int) string {
	year, month, day := time.Now().AddDate(0, 0, day).Date()
	return fmt.Sprintf("%v%02d%02d", year, int(month), day)
}
