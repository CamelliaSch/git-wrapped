# 🦄 Git Wrapped / 代码年报

> 程序员的赛博年度总结 —— 把枯燥的 Git 日志变成可分享的摸鱼报告。
> *Your yearly Git activity, wrapped into a shareable, cyberpunk-style report card.*

[![Go Version](https://img.shields.io/badge/Go-1.20+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/CamelliaSch/git-wrapped/pulls)

[English](sslocal://flow/file_open?url=%23english&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=) | [中文](sslocal://flow/file_open?url=%23chinese&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=)

---

<a id="english"></a>
## English

**Git Wrapped** is an offline CLI tool that dives into your Git history and unearths your coding habits: slacking index, night-owl index, file power couples, commit streaks, and more. It presents the results in a **rainbow terminal UI** and a **cyberpunk HTML card** – all **100% locally**, with zero data leaving your machine.

### ✨ Features

- ☕ **Slacking Index** – Ratio of `fix typo`, `.`, `update` commits.
- 🦉 **Night‑Owl Index** – Percentage of commits made between midnight and 6 AM.
- 📅 **7×52 Heatmap** – ASCII art showing your daily activity over the past year.
- 💕 **File Power Couples** – Which files are always modified together.
- 🔥 **Streak King** – Longest and current commit streaks.
- 🛠️ **Refactor Maniac** – Files you've edited more than 5 times.
- 🎨 **Dual Output**
  - Rainbow terminal report
  - Animated HTML social card
- 📦 **Multi‑repo aggregation**
- 🔒 **100% Offline & Private**


#### Install via Go
```bash
go install github.com/CamelliaSch/git-wrapped@latest
git-wrapped
```

### 📖 Usage
```bash
# Current repo
git-wrapped

# Specific repo
git-wrapped /path/to/repo

# Multi-repo scan
git-wrapped --multi ~/projects --limit 100
```

### 🔒 Privacy
- Zero network requests
- No telemetry
- Read-only Git operations

### 🛠️ Build from Source
```bash
git clone https://github.com/CamelliaSch/git-wrapped.git
cd git-wrapped
go build -o git-wrapped .
```

### 📄 License
[MIT](sslocal://flow/file_open?url=LICENSE&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=)

### 👨‍💻 Author
**CamelliaSch**
- Former Tencent Cloud WAF backend engineer
- B.Eng in Information Security, M.Sc in Computer Science
- GitHub: [@CamelliaSch](sslocal://flow/file_open?url=https%3A%2F%2Fgithub.com%2FCamelliaSch&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=)

---

<a id="chinese"></a>
## 中文

**git-wrapped** 是一个离线命令行工具，深入分析你的 Git 提交记录，生成摸鱼指数、夜猫子指数、文件 CP、连续提交天数等趣味数据，并以**彩虹终端 UI**和**赛博朋克风 HTML 卡片**展示。
**100% 本地运行，绝不上传代码隐私。**

### ✨ 功能特性

- ☕ **摸鱼指数** – 划水类 commit 占比
- 🦉 **夜猫子指数** – 凌晨提交比例
- 📅 **7×52 热力图** – 全年代码活跃度
- 💕 **文件 CP 组合** – 最常一起修改的文件对
- 🔥 **卷王指数** – 最长连续提交天数
- 🛠️ **重构狂魔** – 高频修改文件统计
- 🎨 **双模式输出**：终端彩虹报告 + HTML 分享卡片
- 📦 **多仓库批量分析**
- 🔒 **完全离线，隐私安全**

### 🚀 快速开始

#### 下载二进制文件
1. 前往 [Releases](sslocal://flow/file_open?url=https%3A%2F%2Fgithub.com%2FCamelliaSch%2Fgit-wrapped%2Freleases&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=)
2. 下载对应系统版本
3. 运行：
```bash
./git-wrapped.exe
```

#### Go 安装
```bash
go install github.com/CamelliaSch/git-wrapped@latest
git-wrapped
```

### 📖 使用方法
```bash
# 分析当前仓库
git-wrapped

# 分析指定路径
git-wrapped /path/to/your/repo

# 批量分析多仓库
git-wrapped --multi ~/projects --limit 100
```

### 🔒 隐私承诺
- 零网络请求，完全离线
- 无埋点、无数据上传
- 仅只读读取 Git 日志

### 🛠️ 源码编译
```bash
git clone https://github.com/CamelliaSch/git-wrapped.git
cd git-wrapped
go build -o git-wrapped .
```

### 📄 开源协议
[MIT](sslocal://flow/file_open?url=LICENSE&flow_extra=eyJsaW5rX3R5cGUiOiJjb2RlX2ludGVycHJldGVyIn0=)


## ⭐ Star History

[![Star History Chart](https://api.star-history.com/svg?repos=CamelliaSch/git-wrapped&type=Date)](https://star-history.com/#CamelliaSch/git-wrapped&Date)

如果这个工具让你会心一笑，欢迎点个 Star ⭐ 支持一下！
</sup>


