package response

import (
	"simple-tool/server/internal/models"
)

// ParseRecordListResponse 解析记录列表响应
type ParseRecordListResponse struct {
	Total    int64                `json:"total"`
	PageNo   int                  `json:"page_no"`
	PageSize int                  `json:"page_size"`
	List     []models.ParseRecord `json:"list"`
}
