package payload

import "time"

type JobRequest struct {
	Name        string    `json:"name"`
	Quota       uint      `json:"quota"`
	ExpiryDate  time.Time `json:"expiry_date"`
	JobPosterID uint      `json:"job_poster_id"`
}
