package usecase

import "time"

// データを取得する範囲を生成する
func generateDataRange(now time.Time) (firstDay time.Time, lastDay time.Time) {

	year, month, day := now.Date()

	// 基準となる月初の日
	baseDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	switch {
	case day < 16:
		firstDay = baseDay.AddDate(0, -1, 16)
		lastDay = baseDay.AddDate(0, 0, 15).Add(-1 * time.Second)
	case day >= 16:
		firstDay = baseDay.AddDate(0, 0, 16)
		lastDay = baseDay.AddDate(0, 1, 15).Add(-1 * time.Second)
	}

	return
}

// データを取得する範囲を生成する(課金対象一覧表示用)
func generateDataRangeForShowApp(now time.Time) (firstDay time.Time, lastDay time.Time) {

	year, month, day := now.Date()

	// 基準となる月初の日
	baseDay := time.Date(year, month, 1, 0, 0, 0, 0, time.Local)

	switch {
	case day < 16:
		firstDay = baseDay.AddDate(0, -2, 16)
		lastDay = baseDay.AddDate(0, 0, 15).Add(-1 * time.Second)
	case day >= 16:
		firstDay = baseDay.AddDate(0, 0, 16)
		lastDay = baseDay.AddDate(0, 1, 15).Add(-1 * time.Second)
	}

	return
}
