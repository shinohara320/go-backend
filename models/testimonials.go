package models

type Testimonials struct {
	Id      uint   `json:"id"`
	User    string `json:"user"`
	Message string `json:"message"`
	Image   string `json:"image"`
}
