package responses

import (
	"time"
)

type errorResponse struct {
	Time    time.Duration `json:"duration"`
	Message string        `json:"message"`
}

func Error(start time.Time, message string) errorResponse {
	return errorResponse{
		Time:    time.Duration(time.Since(start).Milliseconds()),
		Message: message,
	}
}

type successResponse struct {
	Status   string        `json:"status"`
	Duration time.Duration `json:"duration"`
	Data     interface{}   `json:"data"`
}

func Success(start time.Time, data interface{}) *successResponse {

	return &successResponse{
		Status:   "success",
		Duration: time.Duration(time.Since(start).Milliseconds()),
		Data:     data,
	}
}
