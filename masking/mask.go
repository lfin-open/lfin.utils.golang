package mask

import "github.com/lfin-open/lfin.utils.golang/strings"

// loopStr 원하는 길이대로 마스킹문자열 생성
// Parameters:
//  - m: 마스킹문자
//  - l: 원하는 길이
// Return:
//  - string: 생성된 masking 문자열
func loopStr(m string, l int) string {
	var mask string
	for i := 1; i <= l; i++ {
		mask += m
	}
	return mask
}

// GenMasking 문자열을 마스킹하여 리턴
// Parameters:
//  - src: 원본
//  - mask: 마스킹할 문자열 (*)
//  - start: 마스킹 시작할 index
//  - length: 마스킹 길이
// Return:
//  - string: 마스킹된 문자열
func GenMasking(src, mask string, start, length int) string {
	result := ""
	if strings.IsEmptyString(src) {
		return result
	}

	l := len(src)
	if start < 0 {
		start = 0
	}
	if start > l || length < 1 {
		return src
	}

	if length > l-start {
		length = l - start
	}

	postSIdx := start + length
	postFIdx := l

	pre := string(src[0:start])
	post := string(src[postSIdx:postFIdx])
	masked := loopStr(mask, length)
	result = pre + masked + post

	return result
}
