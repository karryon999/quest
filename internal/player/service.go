package player

import "time"

// 更新完成任务的天数
func UpdateStreak(p *Player, now time.Time) {
	//第一次完成任务
	if p.LastCompletedDate == nil {
		p.CurrentStreak = 1
	} else if sameDay(*p.LastCompletedDate, now) {
		//今天已经完成过任务，不用修改CurrentStreak
	} else {
		yesterday := now.AddDate(0, 0, -1)
		//如果昨天和上一次完成任务的天是同一天，那就表示连续两天完成任务，所以天数++
		if sameDay(*p.LastCompletedDate, yesterday) {
			p.CurrentStreak++
		} else {
			//如果不一样，那就是中间断了，重新开始记录
			p.CurrentStreak = 1
		}
	}

	if p.CurrentStreak > p.BestStreak {
		p.BestStreak = p.CurrentStreak
	}

	p.LastCompletedDate = &now
}

func sameDay(a, b time.Time) bool {
	return a.Year() == b.Year() &&
		a.Month() == b.Month() &&
		a.Day() == b.Day()
}
