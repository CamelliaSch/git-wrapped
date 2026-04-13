package analytics

import (
	"time"

	"git-wrapped/internal/git"
)

// GetHeatmap 返回过去 52 周、每周 7 天的 Commit 次数矩阵
// 行索引 0-6 对应周日到周六（与 time.Weekday 一致）
func GetHeatmap(commits []git.Commit) [7][52]int {
	var grid [7][52]int

	now := time.Now()
	// 计算 52 周前的起始时间（对齐到周日）
	startDate := now.AddDate(0, 0, -52*7)
	// 调整到当周的周日
	weekday := int(startDate.Weekday())
	startDate = startDate.AddDate(0, 0, -weekday)

	// 遍历所有 Commit，将其放入对应的格子
	for _, c := range commits {
		if c.Date.Before(startDate) {
			continue
		}
		diff := c.Date.Sub(startDate)
		days := int(diff.Hours() / 24)
		if days < 0 || days >= 52*7 {
			continue
		}
		col := days / 7
		row := days % 7
		grid[row][col]++
	}
	return grid
}
