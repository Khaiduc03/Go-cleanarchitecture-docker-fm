package modelCategory



type UpdateCategoryReq struct {
	ID   int    `json:"id" validate:"required,gte=1"`
	Name string `json:"name" validate:"required"`
}

type CreateCategoryReq struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required,oneof=report support"`
}
