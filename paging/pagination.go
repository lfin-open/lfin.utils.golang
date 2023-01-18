package paging

import (
	"math"
)

var DefaultSize = 20

type Pagination struct {
	Offset      int   `json:"offset"`
	CurrentPage int   `json:"currentPage"`
	TotalPages  int   `json:"totalPages"`
	Size        int   `json:"size"`
	Total       int64 `json:"total"`
}

// Calc 페이지 계산하기
func Calc(pageNo, size int, total int64) Pagination {
	p := Pagination{}

	if pageNo <= 0 {
		pageNo = 1
	}
	if size <= 0 {
		size = DefaultSize
	}

	p.Offset = size * (pageNo - 1)
	p.CurrentPage = pageNo
	p.Size = size
	p.Total = total

	// calc total page
	d := float64(total) / float64(size)
	p.TotalPages = int(math.Ceil(d))

	return p
}

//// CalcPaginationFromRequest request 에서 page, pageSize 가져오기
//func CalcPaginationFromRequest(c *gin.Context) Pagination {
//	p := Pagination{}
//
//	page := c.DefaultQuery("page", "1")
//	pageSize := c.DefaultQuery("pageSize", "20")
//
//	p.CurrentPage, _ = strconv.Atoi(page)
//	p.Size, _ = strconv.Atoi(pageSize)
//
//	return p
//}
