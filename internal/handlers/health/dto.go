package health

import "time"

type healthCheckResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
