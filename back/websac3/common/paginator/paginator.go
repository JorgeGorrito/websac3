package paginator

import "fmt"

type paginator[T any] struct {
	page *Page[T]
}

type Page[T any] struct {
	Data         []T   `json:"data"`
	TotalCount   int64 `json:"total_count"`
	Currentpage  uint  `json:"current_page"`
	ItemsPerpage uint  `json:"items_per_page"`
}

type PaginationParams struct {
	Currentpage  uint `json:"current_page" form:"current_page"`
	ItemsPerpage uint `json:"items_per_page" form:"items_per_page"`
}

func New[T any]() *paginator[T] {
	return &paginator[T]{
		page: &Page[T]{
			Data:         []T{},
			TotalCount:   0,
			Currentpage:  0,
			ItemsPerpage: 0,
		},
	}
}

func (p *paginator[T]) CountElements() uint {
	if p.page == nil {
		return 0
	}
	return uint(len(p.page.Data))
}

func (p *paginator[T]) SetData(data []T) *paginator[T] {
	p.page.Data = data
	p.page.ItemsPerpage = p.CountElements()
	return p
}

func (p *paginator[T]) SetTotalCount(totalCount int64) *paginator[T] {
	p.page.TotalCount = totalCount
	return p
}

func (p *paginator[T]) SetCurrentPage(currentPage uint) *paginator[T] {
	p.page.Currentpage = currentPage
	return p
}

func (p *paginator[T]) GetPage() (*Page[T], error) {
	if p.page == nil ||
		p.page.Currentpage == 0 {
		return nil, fmt.Errorf("page is not initialized")
	}
	return p.page, nil
}
