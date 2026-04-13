package main

import (
	"fmt"
	"log"

	"git-wrapped/internal/analytics"
	"git-wrapped/internal/git"
)

func main() {
	fmt.Println("🦄 Git Wrapped - 开始分析你的代码人生...")
	fmt.Println("---------------------------------------")

	// 分析当前目录（就是 git-wrapped 项目本身）
	output, err := git.GetLog(".")
	if err != nil {
		log.Fatal("❌ 获取 Git 日志失败: ", err)
	}

	commits, err := git.ParseLog(output)
	if err != nil {
		log.Fatal("❌ 解析日志失败: ", err)
	}

	if len(commits) == 0 {
		fmt.Println("⚠️  没有找到任何 Commit 记录，是不是还没提交过代码？")
		return
	}

	fmt.Printf("\n📋 共获取到 %d 条 Commit 记录:\n\n", len(commits))
	for i, c := range commits {
		fmt.Printf("%d. [%s] %s: %s\n", i+1, c.Date.Format("01-02 15:04"), c.Author, c.Message)
	}

	// 摸鱼分析
	report := analytics.AnalyzeSlacking(commits)

	fmt.Printf("\n📊 摸鱼报告:\n")
	fmt.Printf("总共 %d 次提交，其中 %d 次疑似在划水（占比 %.1f%%）。\n",
		report.TotalCommits, report.SlackerCommits, report.Ratio)

	if report.Ratio > 30 {
		fmt.Println("💤 鉴定为：职场老油条，typo fix 狂魔。")
	} else if report.Ratio > 10 {
		fmt.Println("🙂 偶尔摸鱼，总体靠谱。")
	} else {
		fmt.Println("⚡ 你是卷王，鉴定完毕。")
	}
}
