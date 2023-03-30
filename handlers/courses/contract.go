package handlers

import "mime/multipart"

type RequestBody struct {
	Name      string         `file:"name"`
	Price     float64        `file:"price"`
	FileImage multipart.File `file:"fileImage"`
}

type ResponseData struct {
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	ImageURL string  `json:"image_url,omitempty"`
	IsFree   bool    `json:"is_free"`
	Rating   float64 `json:"rating,omitempty"`
}
