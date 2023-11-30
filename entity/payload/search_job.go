package payload

type SearchJob struct {
	Name   string `form:"name"`
	Status string `form:"status"`
}
