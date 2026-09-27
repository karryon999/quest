package player

import "testing"

// 编写测试函数
// 测试函数必须以Test开头，go test才会识别
func TestLevel(t *testing.T) {
	// 准备测试数据
	tests := []struct {
		name string
		xp   int
		want int //// want表示我们期望的结果
	}{
		{
			name: "0 XP",
			xp:   0,
			want: 1,
		},
		{
			name: "99 XP",
			xp:   99,
			want: 1,
		},
		{
			name: "100 XP",
			xp:   100,
			want: 2,
		},
		{
			name: "199 XP",
			xp:   199,
			want: 2,
		},
		{
			name: "200 XP",
			xp:   200,
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Player{
				XP: tt.xp,
			}
			got := p.Level() // got表示程序实际算出来的结果
			// 如果结果不一致，这个测试就失败
			if got != tt.want {
				t.Fatalf("Level() = %d, want %d", got, tt.want)
			}
		})
	}
}
