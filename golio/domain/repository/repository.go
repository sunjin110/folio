package repository

type SortType int

const (
	SortTypeAsc  = 1
	SortTypeDesc = 2
)

type Paging struct {
	Offset int
	Limit  int
}
