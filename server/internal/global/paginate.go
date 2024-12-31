package global

import "gorm.io/gorm"

type Pagination struct {
	PageNo   int `form:"page_no" json:"page_no"`
	PageSize int `form:"page_size" json:"page_size"`
}

// Paginate 分页查询
func Paginate(req *Pagination) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		pageNo := req.PageNo
		if pageNo <= 0 {
			pageNo = 1
		}
		pageSize := req.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (pageNo - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
