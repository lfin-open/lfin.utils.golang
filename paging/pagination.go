package paging

import (
	"math"
)

type Pagination struct {
	Offset      int
	CurrentPage int
	MaxPage     int
	Size        int
}

var DefaultSize = 20

// Calc 페이지 계산하기
func Calc(page, size int, total int64) Pagination {
	p := Pagination{}

	if page <= 0 {
		page = 1
	}
	if size <= 0 {
		size = DefaultSize
	}

	p.Offset = size * (page - 1)
	p.CurrentPage = page
	p.Size = size

	// calc maxpage
	d := float64(total) / float64(size)
	p.MaxPage = int(math.Ceil(d))

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
