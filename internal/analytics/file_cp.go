package analytics

import (
	"git-wrapped/internal/git"
)

type FilePair struct {
	FileA string
	FileB string
	Count int
}

// AnalyzeFileCP 分析最常一起修改的文件对 (CP)
func AnalyzeFileCP(repoPath string, commits []git.Commit) []FilePair {
	pairCount := make(map[[2]string]int)

	for _, c := range commits {
		files, err := git.GetChangedFiles(repoPath, c.Hash)
		if err != nil || len(files) < 2 {
			continue
		}
		// 统计所有两两组合
		for i := 0; i < len(files); i++ {
			for j := i + 1; j < len(files); j++ {
				a, b := files[i], files[j]
				if a > b {
					a, b = b, a
				}
				pairCount[[2]string{a, b}]++
			}
		}
	}

	// 找出 Top 3
	type pair struct {
		a, b string
		c    int
	}
	var pairs []pair
	for k, v := range pairCount {
		pairs = append(pairs, pair{k[0], k[1], v})
	}
	// 简单排序取前3
	for i := 0; i < len(pairs)-1; i++ {
		for j := i + 1; j < len(pairs); j++ {
			if pairs[i].c < pairs[j].c {
				pairs[i], pairs[j] = pairs[j], pairs[i]
			}
		}
	}
	topN := 3
	if len(pairs) < topN {
		topN = len(pairs)
	}
	result := make([]FilePair, topN)
	for i := 0; i < topN; i++ {
		result[i] = FilePair{pairs[i].a, pairs[i].b, pairs[i].c}
	}
	return result
}
