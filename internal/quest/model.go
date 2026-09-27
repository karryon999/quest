package quest

import "time"

// Quest记录任务本身的一些参数
type Quest struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	XP        int       `json:"xp"`
	Boss      bool      `json:"boss"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	// *time.Time 表示“指向一个 time.Time 的指针”：值可以为nil，用来判断任务是否完成
	CompletedAt *time.Time `json:"completed_at"`
}
