package git

import (
	"os"
	"path/filepath"
)

// FindRepos 递归查找指定根目录下的所有 Git 仓库
// 返回仓库路径列表
func FindRepos(root string) ([]string, error) {
	var repos []string
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return nil // 忽略权限错误等
		}
		if d.IsDir() && d.Name() == ".git" {
			repos = append(repos, filepath.Dir(path))
			return filepath.SkipDir // 不再进入 .git 内部
		}
		return nil
	})
	return repos, err
}

// AggregateCommits 聚合多个仓库的 Commit 记录
func AggregateCommits(repoPaths []string, limitPerRepo int) ([]Commit, error) {
	var allCommits []Commit

	for _, repoPath := range repoPaths {
		// 获取单个仓库的日志（复用现有 GetLog 和 ParseLog）
		output, err := GetLogWithLimit(repoPath, limitPerRepo)
		if err != nil {
			// 忽略单个仓库的错误，继续处理其他仓库
			continue
		}
		commits, err := ParseLog(output)
		if err != nil {
			continue
		}
		allCommits = append(allCommits, commits...)
	}
	return allCommits, nil
}
