package sql

type Pagination struct {
	Page  int   `json:"page" form:"page"`
	Size  int   `json:"size" form:"size"`
	Total int64 `json:"total"`
}

type Sort struct {
	SortBy string `json:"sortBy" form:"sortBy"`
	Asc    bool   `json:"asc" form:"asc"`
}
