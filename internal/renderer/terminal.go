package renderer

import (
	"fmt"
	"git-wrapped/internal/git"
	"strings"

	"git-wrapped/internal/analytics"
	"github.com/charmbracelet/lipgloss"
)

type TerminalRenderer struct {
	baseStyle   lipgloss.Style
	titleStyle  lipgloss.Style
	labelStyle  lipgloss.Style
	barStyle    lipgloss.Style
	numberStyle lipgloss.Style
}

func NewTerminalRenderer() *TerminalRenderer {
	return &TerminalRenderer{
		baseStyle:   lipgloss.NewStyle().Padding(1, 2).Border(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("63")),
		titleStyle:  lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("213")).MarginBottom(1),
		labelStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color("248")),
		barStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color("51")),
		numberStyle: lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("228")),
	}
}

func (tr *TerminalRenderer) Render(
	commits []git.Commit,
	slacker analytics.SlackerReport,
	nightOwl analytics.NightOwlReport,
	fileCPs []analytics.FilePair,
	streak analytics.StreakReport,
	refactor analytics.RefactorReport,
) {
	title := tr.titleStyle.Render("🦄 Git Wrapped 2025")

	slackerBar := tr.generateBar(slacker.Ratio, 20)
	nightOwlBar := tr.generateBar(nightOwl.Ratio, 20)

	// 基础内容：摸鱼指数 + 夜猫子指数 + 评语
	baseContent := lipgloss.JoinVertical(lipgloss.Left,
		tr.labelStyle.Render("📊 摸鱼指数: ")+tr.numberStyle.Render(fmt.Sprintf("%.1f%%", slacker.Ratio))+"  "+slackerBar,
		tr.labelStyle.Render("🌙 夜猫子指数: ")+tr.numberStyle.Render(fmt.Sprintf("%.1f%%", nightOwl.Ratio))+"  "+nightOwlBar,
		"",
		tr.getVerdict(slacker.Ratio, nightOwl.Ratio),
	)

	// 热力图
	heatmapGrid := analytics.GetHeatmap(commits)
	heatmap := tr.renderHeatmap(heatmapGrid)

	// 文件 CP 组合
	cpSection := tr.renderFileCPs(fileCPs)

	// 更多标签
	extraSection := tr.renderExtraLabels(streak, refactor)

	// 将所有部分垂直拼接（只声明一次 fullContent）
	fullContent := lipgloss.JoinVertical(lipgloss.Left,
		baseContent,
		heatmap,
		cpSection,
		extraSection,
	)

	// 最终渲染
	fmt.Println(tr.baseStyle.Render(title + "\n\n" + fullContent))
}

func (tr *TerminalRenderer) renderFileCPs(pairs []analytics.FilePair) string {
	if len(pairs) == 0 {
		return ""
	}
	var sb strings.Builder
	sb.WriteString("\n" + tr.labelStyle.Render("💕 最强文件 CP (一起修改次数)"))
	for _, p := range pairs {
		sb.WriteString(fmt.Sprintf("\n%s & %s: %d 次", p.FileA, p.FileB, p.Count))
	}
	return sb.String()
}

func (tr *TerminalRenderer) generateBar(percent float64, width int) string {
	filled := int(percent / 100 * float64(width))
	if filled > width {
		filled = width
	}
	empty := width - filled
	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	return tr.barStyle.Render(bar)
}

func (tr *TerminalRenderer) getVerdict(slackerRatio, nightOwlRatio float64) string {
	var verdict strings.Builder
	if slackerRatio > 30 {
		verdict.WriteString("💤 你是职场皮划艇冠军，typofix 艺术家。\n")
	} else {
		verdict.WriteString("⚡ 代码机器，卷王本王。\n")
	}
	if nightOwlRatio > 20 {
		verdict.WriteString("🦉 月亮不睡你不睡，你是秃头小宝贝。")
	} else {
		verdict.WriteString("🌞 作息健康，发量惊人。")
	}
	return tr.labelStyle.Render(verdict.String())
}

// renderHeatmap 方法
func (tr *TerminalRenderer) renderHeatmap(grid [7][52]int) string {
	var builder strings.Builder
	builder.WriteString("\n" + tr.labelStyle.Render("📅 过去一年代码活跃度 (7x52)") + "\n")

	days := []string{"日", "一", "二", "三", "四", "五", "六"}
	for row := 0; row < 7; row++ {
		builder.WriteString(days[row] + " ")
		for col := 0; col < 52; col++ {
			count := grid[row][col]
			builder.WriteString(tr.getHeatChar(count))
		}
		builder.WriteString("\n")
	}
	builder.WriteString("   少 " + tr.getHeatChar(0) + tr.getHeatChar(3) + tr.getHeatChar(6) + tr.getHeatChar(10) + tr.getHeatChar(15) + " 多\n")
	return builder.String()
}

func (tr *TerminalRenderer) getHeatChar(count int) string {
	switch {
	case count == 0:
		return "░"
	case count < 3:
		return "▒"
	case count < 6:
		return "▓"
	default:
		return "█"
	}
}

func (tr *TerminalRenderer) renderExtraLabels(streak analytics.StreakReport, refactor analytics.RefactorReport) string {
	var sb strings.Builder
	sb.WriteString("\n" + tr.labelStyle.Render("🏆 更多标签"))

	// 卷王指数
	sb.WriteString(fmt.Sprintf("\n🔥 最长连续提交: %d 天", streak.LongestStreak))
	if streak.CurrentStreak > 0 {
		sb.WriteString(fmt.Sprintf(" (当前已连续 %d 天)", streak.CurrentStreak))
	}

	// 重构狂魔
	if refactor.FileCount > 0 {
		sb.WriteString(fmt.Sprintf("\n🛠️ 重构狂魔: 修改超过 %d 次的文件有 %d 个", refactor.Threshold, refactor.FileCount))
		// 显示 Top 1 文件名
		maxFile := ""
		maxCount := 0
		for f, c := range refactor.FileDetails {
			if c > maxCount {
				maxCount = c
				maxFile = f
			}
		}
		if maxFile != "" {
			sb.WriteString(fmt.Sprintf(" (冠军: %s, %d 次)", maxFile, maxCount))
		}
	} else {
		sb.WriteString("\n✨ 代码洁癖: 没有文件被反复蹂躏")
	}

	return sb.String()
}
