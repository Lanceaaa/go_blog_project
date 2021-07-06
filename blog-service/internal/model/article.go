package model

type Article struct {
	*Model
	Title string `json:"title"`
	desc string `json:"desc"`
	content string `json:"content"`
	coverImageUrl string `json:"cover_image_url"`
	state uint8 `json:"state"`
}