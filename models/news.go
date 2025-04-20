package models

type News struct {
	BaseModel
	Img      string `json:"img"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Markdown string `json:"markdown"`
}
