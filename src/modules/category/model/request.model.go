package modelCategory

type UpdateCategoryReq struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type CreateCategoryReq struct {
	Name       string `json:"name"`
}
