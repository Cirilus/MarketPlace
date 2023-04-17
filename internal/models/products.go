package models

type Product struct {
	Id          int     `json:"id,omitempty"`
	Title       string  `json:"title,omitempty"`
	Cost        int     `json:"cost,omitempty"`
	Description string  `json:"description,omitempty"`
	Author      User    `json:"author"`
	Category    string  `json:"category,omitempty"`
	Rate        float32 `json:"rate,omitempty"`
}
