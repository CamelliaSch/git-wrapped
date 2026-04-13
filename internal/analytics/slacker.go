package analytics

import (
	"strings"

	"git-wrapped/internal/git"
)

// SlackerReport 摸鱼分析报告
type SlackerReport struct {
	TotalCommits   int
	SlackerCommits int
	Ratio          float64
}

// AnalyzeSlacking 分析 Commit 中的摸鱼迹象
func AnalyzeSlacking(commits []git.Commit) SlackerReport {
	report := SlackerReport{TotalCommits: len(commits)}

	for _, c := range commits {
		msg := strings.ToLower(c.Message)

		// 摸鱼关键词字典（后续可以继续丰富）
		isSlacking := strings.Contains(msg, "typo") ||
			strings.Contains(msg, "minor") ||
			strings.Contains(msg, "update") ||
			strings.Contains(msg, "wip") ||
			strings.Contains(msg, "chore") ||
			strings.Contains(msg, "fix typo") ||
			msg == "." ||
			msg == "update"

		if isSlacking {
			report.SlackerCommits++
		}
	}

	if report.TotalCommits > 0 {
		report.Ratio = float64(report.SlackerCommits) / float64(report.TotalCommits) * 100
	}

	return report
}
