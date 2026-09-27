// quest负责Quest相关的业务规则
package quest

import "time"

// 根据给定的标题、XP 和 Boss 状态创建一个新任务
func New(title string, xp int, boss bool) Quest {
	q := Quest{
		Title:       title,
		XP:          xp,
		Boss:        boss,
		Completed:   false,
		CreatedAt:   time.Now(),
		CompletedAt: nil,
	}
	return q
}

// 获取新建任务的ID：已有任务中最大的ID值+1
func NextID(quests []Quest) int {
	maxID := 0
	for _, q := range quests {
		if q.ID > maxID {
			maxID = q.ID
		}
	}
	return maxID + 1
}
