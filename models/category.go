package models

type Category struct {
	BaseModel
	ParentCategoryID *uint64   `json:"parentCategoryID"`
	ParentCategory   *Category `json:"parentCategory"`
	Name             string    `json:"name"`
}
