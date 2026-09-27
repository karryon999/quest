package player

import "time"

// Player记录的是“某个玩家目前的游戏进度”。

type Player struct {
	XP int `json:"xp"`
	//CurrentStreak: 当前连续完成任务的天数
	CurrentStreak int `json:"current_streak"`
	//BestStreak:历史最高连续天数
	BestStreak        int        `json:"best_streak"`
	LastCompletedDate *time.Time `json:"last_completed_date"`
}

// Level()是一个方法，不是函数，因为方法才有接收者(p Player)
// 接收者表述Level()是Player的方法
// 根据XP计算出所处的level：
// 0 ~ 99 XP      → Level 1
// 100 ~ 199 XP   → Level 2
// 200 ~ 299 XP   → Level 3

func (p Player) Level() int {
	return p.XP/100 + 1
}
