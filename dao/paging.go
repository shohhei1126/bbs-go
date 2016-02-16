package dao

type Paging struct {
	Limit   uint64
	Offset  uint64
	OrderBy string
}