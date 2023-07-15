package model

type MenuItem struct {
	ID         int         `json:"id"`
	ParentID   int         `json:"parent_id"`
	Title      string      `json:"title"`
	Name       string      `json:"name"`
	Route      string      `json:"route"`
	Icon       string      `json:"icon"`
	IsChildren bool        `json:"is_children"`
	Children   []*MenuItem `json:"children"`
}
