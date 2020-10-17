package pagination

type pagination struct {
	PageSize    int `json:"page_size"`
	CurrentPage int `json:"current_page"`
	PageTotal   int `json:"page_total"`
	DataTotal   int `json:"data_total"`
}

func New(pageSize, currentPage, dataTotal int) *pagination {
	pageTotal := dataTotal / pageSize
	if dataTotal%pageSize != 0 {
		pageTotal += 1
	}
	return &pagination{
		PageSize:    pageSize,
		CurrentPage: currentPage,
		DataTotal:   dataTotal,
		PageTotal:   pageTotal,
	}
}

func (page *pagination) GetPageOffset() int {
	return (page.CurrentPage - 1) * page.PageSize
}
