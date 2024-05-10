package entity

type Pagination struct {
	Query *string `form:"query"`
	Limit int     `form:"limit"`
	Page  int     `form:"page"`
}
