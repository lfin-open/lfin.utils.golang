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

package date

import "time"

// GetCurrDateTimeYMD24HMMSSdot6Z
// 현재시간 문자열 리턴
// 포멧: YYYY-MM-DD HH:MM:SS.ssssss MST
// eg) 2022-01-07 14:56:44.545722 KST
func GetCurrDateTimeYMD24HMMSSdot6Z() string {
	// 현재시간 문자열을 리턴
	// 포멧: YYYY-MM-DD HH:MM:SS.ssssss MST
	// eg) 2022-01-07 14:56:44.545722 KST
	return time.Now().Format("2006-01-02 15:04:05.000000 MST")
}

// ConvertTimeYMD24HMMSSdot6Z
// time.Time 을 문자열로 변환
// 포멧: YYYY-MM-DD HH:MM:SS.ssssss MST
// eg) 2022-01-07 14:56:44.545722 KST
func ConvertTimeYMD24HMMSSdot6Z(time time.Time) string {
	// 현재시간 문자열을 리턴
	// 포멧: YYYY-MM-DD HH:MM:SS.ssssss MST
	// eg) 2022-01-07 14:56:44.545722 KST
	return time.Format("2006-01-02 15:04:05.000000 MST")
}

// GetCurrDateTimeYMD24HMMSS
// 현재시간 문자열 리턴
// 포멧: YYYY-MM-DD HH:MM:SS
// eg) 2022-01-07 14:56:44
func GetCurrDateTimeYMD24HMMSS() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// GetCurrDateTimeYMD24HMMSSNoDiv 현재시간 문자열 리턴 (구문자 없이 붙여서)
// 포멧: YYYYMMDDHHMMSS
// eg) 20220107145644
func GetCurrDateTimeYMD24HMMSSNoDiv() string {
	return time.Now().Format("20060102150405")
}

// GetCurrentUnixTimestampSec 현재시간의 초단위 unixtimestamp
func GetCurrentUnixTimestampSec() int64 {
	return time.Now().Unix()
}

// GetCurrentUnixTimestampNano 현재시간의 나노초 단위 unixtimestamp
func GetCurrentUnixTimestampNano() int64 {
	return time.Now().UnixNano()
}

// GetCurrentUnixTimestampMill 현재시간의 밀리초 단위 unixtimestamp
func GetCurrentUnixTimestampMill() int64 {
	return time.Now().UnixMilli()
}
