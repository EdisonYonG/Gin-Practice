package models

type Nav struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Status int    `json:"status"`
	Url    string `json:"url"`
	Sort   int    `json:"sort"`
}

func (Nav) TableName() string {
	return "nav"
}
