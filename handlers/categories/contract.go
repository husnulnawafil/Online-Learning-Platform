package handlers

type ResponseData struct {
	Name           string  `json:"name"`
	TotalSubcriber int64   `json:"total_subcriber,omitempty"`
	AverageRating  float32 `json:"average_rating,omitempty"`
}
