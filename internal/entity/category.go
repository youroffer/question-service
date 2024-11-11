package entity

import (
	api "github.com/himmel520/question-service/api/oas"
	"github.com/himmel520/question-service/internal/lib/convert"
)

type Category struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Public bool   `json:"public"`
}

func CategoryToApi(c *Category) *api.Category {
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

type CategoriesResp struct {
	Data    []*Category `json:"data"`
	Page    int        `json:"page"`
	Pages   int        `json:"pages"`
	PerPage int        `json:"per_page"`
}

func (c *CategoriesResp) CategoriesRespToApi() *api.CategoriesResp {
	return &api.CategoriesResp{
		Data: convert.MapSlice(c.Data, CategoryToApi),
		Page: c.Page,
		Pages: c.Pages,
		PerPage: c.PerPage,
	}
}
