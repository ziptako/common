package utils

// PaginationParams 分页参数结构
type PaginationParams struct {
	Page     int32 // 页码（从1开始）
	PageSize int32 // 每页数量
}

// PaginationResult 分页计算结果
type PaginationResult struct {
	Limit  int32 // 查询限制数量
	Offset int32 // 查询偏移量
}

// CalculatePagination 计算分页参数
// 输入页码和每页数量，返回数据库查询所需的limit和offset
func CalculatePagination(page, pageSize int32) PaginationResult {
	// 参数验证和默认值设置
	if page <= 0 {
		page = 1 // 默认第1页
	}
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}
	if pageSize > 100 {
		pageSize = 100 // 最大每页100条
	}

	// 计算offset
	offset := (page - 1) * pageSize

	return PaginationResult{
		Limit:  pageSize,
		Offset: offset,
	}
}

// ValidatePaginationParams 验证分页参数
// 返回验证后的安全参数
func ValidatePaginationParams(params PaginationParams) PaginationParams {
	if params.Page <= 0 {
		params.Page = 1
	}
	if params.PageSize <= 0 {
		params.PageSize = 10
	}
	if params.PageSize > 100 {
		params.PageSize = 100
	}
	return params
}
