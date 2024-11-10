package entity

import api "github.com/himmel520/question-service/api/oas"

type Category struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Public bool   `json:"public"`
}

func (c *Category) CategoryToApi() *api.Category {
	return &api.Category{
		ID:     c.ID,
		Title:  c.Title,
		Public: c.Public,
	}
}

type CategoryUpdate struct {
	Title  Optional[string] `json:"title"`
	Public Optional[bool]   `json:"public"`
}

func (c *CategoryUpdate) IsSet() bool {
	return c.Title.Set || c.Public.Set
}
