package web

type CategoryUpdateRequest struct {
	Id   int64  `validate:"required"`
	Name string `validate:"required,min=1,max=100" json:"name"`
}
