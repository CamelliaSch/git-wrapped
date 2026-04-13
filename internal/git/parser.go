package git

import (
	"strings"
	"time"
)

// Commit 表示一条 Git 提交记录
type Commit struct {
	Hash    string
	Author  string
	Date    time.Time
	Message string
}

// ParseLog 将 git log 的原始字符串解析为 Commit 切片
func ParseLog(logOutput string) ([]Commit, error) {
	lines := strings.Split(strings.TrimSpace(logOutput), "\n")
	commits := make([]Commit, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}
		// 按 | 分割，最多分 4 部分（因为 message 中可能包含 |）
		parts := strings.SplitN(line, "|", 4)
		if len(parts) < 4 {
			continue
		}

		// 解析时间
		t, err := time.Parse(time.RFC3339, parts[2])
		if err != nil {
			// 时间解析失败，使用零值占位
			t = time.Time{}
		}

		commits = append(commits, Commit{
			Hash:    parts[0],
			Author:  parts[1],
			Date:    t,
			Message: parts[3],
		})
	}

	return commits, nil
}
