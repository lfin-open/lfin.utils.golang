/*
 * Copyright (c) 2022 LFin and others.
 *
 * All rights reserved.
 *
 *
 * Contributors:
 *    Ted KIM
 *
 */

package strings

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsEmptyString(t *testing.T) {
	var src string // src = ""
	assert.Equal(t, true, IsEmptyString(src), fmt.Sprintf("[%s] -> Empty", src))

	src = "NoEmpty"
	assert.Equal(t, false, IsEmptyString(src), fmt.Sprintf("[%s] -> NoEmpty", src))

	src = "   "
	assert.Equal(t, true, IsEmptyString(src), fmt.Sprintf("[%s] -> Empty", src))
}

func TestI64ToS(t *testing.T) {
	var i int64 = 1651027224746
	s := I64ToS(i)
	t.Logf("int64:%d --> string:%s", i, s)
}

func TestToSnakeCase(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"TestToSnakeCase_1", args{"TestToSnakeCase"}, "test_to_snake_case"},
		{"TestToSnakeCase_2", args{"testtosnakeCase"}, "testtosnake_case"},
		{"TestToSnakeCase_3", args{"testtosnakecase"}, "testtosnakecase"},
		{"TestToSnakeCase_4", args{"Test To Snake Case"}, "test_to_snake_case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, ToSnakeCase(tt.args.str), "ToSnakeCase(%v)", tt.args.str)
		})
	}
}
