package paging

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalc(t *testing.T) {
	page := Pagination{Size: 25, CurrentPage: 5, MaxPage: 5, Offset: 100}
	assert.Equal(t, page, Calc(5, 25, 104))

	page2 := Pagination{Size: 10, CurrentPage: 7, MaxPage: 11, Offset: 60}
	assert.Equal(t, page2, Calc(7, 10, 105))

	page0 := Pagination{Size: 10, CurrentPage: 1, MaxPage: 1, Offset: 0}
	assert.Equal(t, page0, Calc(1, 10, 5))
}
