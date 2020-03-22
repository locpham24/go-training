package form

type Note struct {
	Title      string `json:"title"`
	Completed  bool   `json:"completed"`
	CategoryId int    `json:"category_id"`
}
