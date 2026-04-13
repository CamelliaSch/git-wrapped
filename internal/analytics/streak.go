package analytics

import (
	"sort"
	"time"

	"git-wrapped/internal/git"
)

// StreakReport 连续提交分析
type StreakReport struct {
	LongestStreak int // 最长连续提交天数
	CurrentStreak int // 当前连续提交天数（从今天往回算）
}

// AnalyzeStreak 分析连续提交天数
func AnalyzeStreak(commits []git.Commit) StreakReport {
	if len(commits) == 0 {
		return StreakReport{}
	}

	// 提取所有提交日期（按天去重）
	dateSet := make(map[string]bool)
	for _, c := range commits {
		dateStr := c.Date.Format("2006-01-02")
		dateSet[dateStr] = true
	}

	// 转为排序后的日期列表
	var dates []time.Time
	for d := range dateSet {
		t, _ := time.Parse("2006-01-02", d)
		dates = append(dates, t)
	}
	sort.Slice(dates, func(i, j int) bool { return dates[i].Before(dates[j]) })

	// 计算最长连续天数
	longest := 0
	current := 1
	for i := 1; i < len(dates); i++ {
		diff := dates[i].Sub(dates[i-1]).Hours() / 24
		if diff == 1 {
			current++
		} else {
			if current > longest {
				longest = current
			}
			current = 1
		}
	}
	if current > longest {
		longest = current
	}

	// 计算当前连续天数（从今天往回数）
	today := time.Now().Truncate(24 * time.Hour)
	currentStreak := 0
	for i := len(dates) - 1; i >= 0; i-- {
		expectedDate := today.AddDate(0, 0, -currentStreak)
		if dates[i].Equal(expectedDate) {
			currentStreak++
		} else {
			break
		}
	}

	return StreakReport{
		LongestStreak: longest,
		CurrentStreak: currentStreak,
	}
}
