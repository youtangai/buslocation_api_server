package model

type RequestList struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

type RequestInfo struct {
	StartID string `json:"start_id"`
	EndID   string `json:"end_id"`
}
