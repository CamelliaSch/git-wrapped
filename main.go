package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"runtime"

	"git-wrapped/internal/analytics"
	"git-wrapped/internal/git"
	"git-wrapped/internal/renderer"
)

func main() {
	// 定义参数
	multiDir := flag.String("multi", "", "分析指定目录下的所有 Git 仓库（聚合模式）")
	limit := flag.Int("limit", 200, "每个仓库最多分析的 Commit 数量")
	flag.Parse()

	var commits []git.Commit
	var repoPath string

	if *multiDir != "" {
		// 多仓库模式
		fmt.Printf("🔍 正在扫描目录: %s\n", *multiDir)
		repos, err := git.FindRepos(*multiDir)
		if err != nil {
			log.Fatal("扫描仓库失败: ", err)
		}
		if len(repos) == 0 {
			log.Fatal("未找到任何 Git 仓库")
		}
		fmt.Printf("📦 发现 %d 个仓库\n", len(repos))
		for _, r := range repos {
			fmt.Printf("   - %s\n", filepath.Base(r))
		}

		commits, err = git.AggregateCommits(repos, *limit)
		if err != nil {
			log.Fatal("聚合 Commit 失败: ", err)
		}
		repoPath = *multiDir // 用于后续文件分析（多仓库下文件分析可能意义不大，可选择性跳过）
	} else {
		// 单仓库模式（默认当前目录）
		repoPath = "."
		if flag.NArg() > 0 {
			repoPath = flag.Arg(0)
		}
		output, err := git.GetLogWithLimit(repoPath, *limit)
		if err != nil {
			log.Fatal("获取 Git 日志失败: ", err)
		}
		commits, err = git.ParseLog(output)
		if err != nil {
			log.Fatal("解析日志失败: ", err)
		}
	}

	if len(commits) == 0 {
		fmt.Println("⚠️ 没有找到任何 Commit 记录")
		return
	}

	// 后续分析与单仓库模式完全一致
	slackerReport := analytics.AnalyzeSlacking(commits)
	nightOwlReport := analytics.AnalyzeNightOwl(commits)
	fileCPs := analytics.AnalyzeFileCP(repoPath, commits) // 多仓库下文件分析可能不准确，但保留
	streakReport := analytics.AnalyzeStreak(commits)
	refactorReport := analytics.AnalyzeRefactor(commits, repoPath, 5)

	terminalRenderer := renderer.NewTerminalRenderer()
	terminalRenderer.Render(commits, slackerReport, nightOwlReport, fileCPs, streakReport, refactorReport)
	htmlPath := "git-wrapped-report.html"
	// 生成 HTML
	err := renderer.GenerateHTMLCard(slackerReport, nightOwlReport, "git-wrapped-report.html")
	if err != nil {
		log.Printf("⚠️ HTML 生成失败: %v", err)
	} else {
		fmt.Println("\n📸 HTML 报告已生成: git-wrapped-report.html")
		openBrowser(htmlPath)
	}
}

func openBrowser(path string) {
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", "", path)
	case "darwin":
		cmd = exec.Command("open", path)
	case "linux":
		cmd = exec.Command("xdg-open", path)
	default:
		fmt.Printf("⚠️  请手动打开报告: %s\n", path)
		return
	}
	if err := cmd.Start(); err != nil {
		fmt.Printf("⚠️  自动打开浏览器失败: %v\n", err)
	}
}
