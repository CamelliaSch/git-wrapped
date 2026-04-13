package git

import (
	"bytes"
	"fmt"
	"os/exec"
)

// GetLog 获取指定路径仓库的 Git 日志（最近 50 条，便于测试）
// path: 仓库路径，通常传 "." 表示当前目录
func GetLog(path string) (string, error) {
	// 1. 检查 git 命令是否可用
	_, err := exec.LookPath("git")
	if err != nil {
		return "", fmt.Errorf("未找到 git 命令，请确认已安装 Git 并添加到 PATH")
	}

	// 2. 构建安全命令参数（不拼接字符串，防止注入）
	args := []string{
		"-C", path,
		"log",
		"--pretty=format:%H|%an|%ad|%s",
		"--date=iso-strict",
		"--no-merges",
		"-n", "50", // 获取最近 50 条，初期调试足够
	}

	cmd := exec.Command("git", args...)
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git 命令执行失败: %w", err)
	}

	return out.String(), nil
}
