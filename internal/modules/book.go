package modules

type Book struct {
	ID     uint   `json:"id"`
	Title  string `json: "title"`
	Author string `json: "author"`
	Genre  string `json: "genre"`
}
