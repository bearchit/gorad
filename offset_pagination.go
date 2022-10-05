package gorad

type OffsetPageInfo struct {
	TotalCount int `json:"totalCount"`
	Offset     int `json:"offset"`
	Limit      int `json:"limit"`
}

type OffsetConnection[T any] struct {
	TotalCount int            `json:"totalCount"`
	Edges      []*T           `json:"edges"`
	PageInfo   OffsetPageInfo `json:"pageInfo"`
}
