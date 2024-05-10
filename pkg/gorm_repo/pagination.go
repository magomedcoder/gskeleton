package gorm_repo

import (
	"fmt"

	"gorm.io/gorm"
)

type Pagination struct {
	Page     int
	PageSize int
}

func (p *Pagination) GetOffset() int {
	return (p.GetPage() - 1) * p.GetLimit()
}

func (p *Pagination) GetPage() int {
	if p.Page == 0 {
		p.Page = 1
	}
	return p.Page
}

func (p *Pagination) GetLimit() int {
	if p.PageSize == 0 {
		p.PageSize = 15
	}
	return p.PageSize
}

func Paginate(pagination *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}

func CustomPagination(pagination *Pagination) string {
	return fmt.Sprintf("LIMIT %d OFFSET %d", pagination.GetLimit(), pagination.GetOffset())
}
