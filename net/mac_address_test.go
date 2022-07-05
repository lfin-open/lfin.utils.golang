package net

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToMacAddressStyle(t *testing.T) {
	s := "80CA4B4B6F0B"
	exp := "80:CA:4B:4B:6F:0B"
	result := ConvertToMacAddressStyle(s, ":")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80CA4B4B6F0B"
	exp = "80-CA-4B-4B-6F-0B"
	result = ConvertToMacAddressStyle(s, "-")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80CA4B4B6F0B"
	exp = "80::CA::4B::4B::6F::0B"
	result = ConvertToMacAddressStyle(s, "::")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80CA4B4B6F0"
	exp = "80:CA:4B:4B:6F:0"
	result = ConvertToMacAddressStyle(s, ":")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80CA4B4B6F0"
	exp = "80CA4B4B6F0"
	result = ConvertToMacAddressStyle(s, "")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80CA4B4B6F0"
	exp = "80:CA:4B:4B:6F:0"
	result = ConvertToMacAddressStyle(s, ":")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80"
	exp = "80"
	result = ConvertToMacAddressStyle(s, ":")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = "80E"
	exp = "80:E"
	result = ConvertToMacAddressStyle(s, ":")
	t.Logf("%s --> %s", s, result)
	assert.Equal(t, exp, result)

	s = ""
	exp = ""
	result = ConvertToMacAddressStyle(s, ":")
	t.Logf("[%s] --> [%s]", s, result)
	assert.Equal(t, exp, result)

}
