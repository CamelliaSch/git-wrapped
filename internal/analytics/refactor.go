package analytics

import (
	"git-wrapped/internal/git"
)

// RefactorReport 重构狂魔分析
type RefactorReport struct {
	FileCount   int            // 被频繁修改的文件数量（超过阈值）
	FileDetails map[string]int // 每个文件的修改次数
	Threshold   int            // 阈值
}

// AnalyzeRefactor 统计修改次数超过阈值的文件
func AnalyzeRefactor(commits []git.Commit, repoPath string, threshold int) RefactorReport {
	fileCounts := make(map[string]int)

	for _, c := range commits {
		files, err := git.GetChangedFiles(repoPath, c.Hash)
		if err != nil {
			continue
		}
		for _, f := range files {
			fileCounts[f]++
		}
	}

	// 筛选超过阈值的文件
	details := make(map[string]int)
	frequentCount := 0
	for file, count := range fileCounts {
		if count >= threshold {
			frequentCount++
			details[file] = count
		}
	}

	return RefactorReport{
		FileCount:   frequentCount,
		FileDetails: details,
		Threshold:   threshold,
	}
}
