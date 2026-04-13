package analytics

import (
	"git-wrapped/internal/git"
)

type NightOwlReport struct {
	TotalCommits int
	NightCommits int
	Ratio        float64
}

func AnalyzeNightOwl(commits []git.Commit) NightOwlReport {
	report := NightOwlReport{TotalCommits: len(commits)}
	for _, c := range commits {
		hour := c.Date.Hour()
		// 凌晨 0-5 点视为夜猫子
		if hour >= 0 && hour < 6 {
			report.NightCommits++
		}
	}
	if report.TotalCommits > 0 {
		report.Ratio = float64(report.NightCommits) / float64(report.TotalCommits) * 100
	}
	return report
}
