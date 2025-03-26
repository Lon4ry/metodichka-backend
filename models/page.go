package models

type Page struct {
	BaseModel
	Markdown   string   `json:"markdown"`
	Category   Category `json:"category"`
	CategoryID uint64   `json:"categoryID" gorm:"unique"`
}
