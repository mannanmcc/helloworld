package handlers

type apiResponse struct {
	Source      string  `json:"source"`
	Destination string  `json:"destination"`
	Rate        float64 `json:"rate"`
}
