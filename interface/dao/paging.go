package dao

type Paging struct {
	Limit   uint
	Offset  uint
	OrderBy string
}