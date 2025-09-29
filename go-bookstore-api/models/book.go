package models

type Book struct {
	ID          uint     `gorm:"primaryKey" json:"id"`
	Title       string   `gorm:"not null" json:"title"`
	Description string   `json:"description"`
	ImageURL    string   `json:"image_url"`
	ReleaseYear int      `json:"release_year"`
	Price       int      `json:"price"`
	TotalPage   int      `json:"total_page"`
	Thickness   string   `json:"thickness"`
	CategoryID  uint     `json:"category_id"`
	Category    Category `json:"category"`
}