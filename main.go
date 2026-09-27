package main

import (
	"os"

	"github.com/karryon999/quest/cmd"
)

// quest add       ✅ 创建任务并持久化
// quest list      ✅ 查看任务
// quest delete    ✅ 删除任务
// quest done      ✅ 完成任务
//
//	├── Completed
//	├── CompletedAt
//	├── XP
//	└── Streak
//
// quest stats     ✅ Level / XP / Completed / Streak
//
// quests.json     ✅
// player.json     ✅
func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
