package payload

type ApplicationRequest struct {
	UserID uint `json:"user_id"`
	JobID  uint `json:"job_id"`
}

type ApplicationResponse struct {
}
