package paginator

import "fmt"

type paginator[T any] struct {
	page *page[T]
}

type page[T any] struct {
	Data         []T
	TotalCount   uint
	Currentpage  uint
	ItemsPerpage uint
}

type PaginationParams struct {
	Currentpage  uint
	ItemsPerpage uint
}

func New[T any]() *paginator[T] {
	return &paginator[T]{
		page: &page[T]{
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

func (p *paginator[T]) SetTotalCount(totalCount uint) *paginator[T] {
	p.page.TotalCount = totalCount
	return p
}

func (p *paginator[T]) SetCurrentPage(currentPage uint) *paginator[T] {
	p.page.Currentpage = currentPage
	return p
}

func (p *paginator[T]) GetPage() (*page[T], error) {
	if p.page == nil ||
		p.page.Data == nil ||
		p.page.Currentpage == 0 ||
		p.page.ItemsPerpage == 0 {
		return nil, fmt.Errorf("page is not initialized")
	}
	return p.page, nil
}
