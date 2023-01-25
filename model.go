package sql

type Pagination struct {
	Page  int   `form:"page"`
	Size  int   `form:"size"`
}

type Sort struct {
	SortBy string `form:"sortBy"`
	Asc    bool   `form:"asc"`
}
