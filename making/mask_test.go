package mask

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenMasking(t *testing.T) {

	src := ""
	exp := ""
	result := GenMasking(src, "*", 2, 3)
	t.Logf("GenMasking src:%s expected:%s result:%s", src, exp, result)
	assert.Equal(t, exp, result)

	src = "01234567890"
	exp = "01*****7890"
	result = GenMasking(src, "*", 2, 5)
	t.Logf("GenMasking src:%s expected:%s result:%s", src, exp, result)
	assert.Equal(t, exp, result)

	src = "01234567890"
	exp = "***********"
	result = GenMasking(src, "*", -1, 20)
	t.Logf("GenMasking src:%s expected:%s result:%s", src, exp, result)
	assert.Equal(t, exp, result)

	src = "01234567890"
	exp = "******67890"
	result = GenMasking(src, "*", 0, 6)
	t.Logf("GenMasking src:%s expected:%s result:%s", src, exp, result)
	assert.Equal(t, exp, result)

	src = "4e0b15f0-c9ac-11ec-be25-a978960d4284"
	exp = "4e0b15f0-c9ac-**********************"
	result = GenMasking(src, "*", 14, 100)
	t.Logf("GenMasking \nsrc     :%s \nexpected:%s \nresult  :%s", src, exp, result)
	assert.Equal(t, exp, result)

	src = "D10F5178FDFB9824F5641651CC44B3A5E7BF2A145FF88594C1BC207DF6CC9F77CFE09D53F7C9206C1A838BE1FA737ABE71FAA26FC0EDB7C6187A4D136BFA62E8"
	exp = "D10F5178FDFB9824F564************************************************************************************************************"
	result = GenMasking(src, "*", 20, 1000)
	t.Logf("GenMasking \nsrc     :%s \nexpected:%s \nresult  :%s", src, exp, result)
	assert.Equal(t, exp, result)

}
