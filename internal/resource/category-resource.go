package resource

import "github.com/Fajar3108/online-course-be/pkg/model"

type CategoryResource struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Slug      string `json:"slug"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewCategoryResource(category *model.Category) *CategoryResource {
	return &CategoryResource{
		ID:        category.ID,
		Name:      category.Name,
		Slug:      category.Slug,
		CreatedAt: category.CreatedAt.Time.Format("2006-01-02 15:04:05"),
		UpdatedAt: category.UpdatedAt.Time.Format("2006-01-02 15:04:05"),
	}
}
