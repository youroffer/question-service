package entity

import api "github.com/himmel520/question-service/api/oas"

type Category struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

func(c *Category) CategoryToApi() *api.Category {
	return &api.Category{
		ID: c.ID,
		Title: c.Title,
	}
}
