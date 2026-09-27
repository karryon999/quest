package quest

import "testing"

func TestNextID(t *testing.T) {
	tests := []struct {
		name   string
		quests []Quest
		want   int
	}{
		{
			name:   "no quests",
			quests: []Quest{},
			want:   1,
		},
		{
			name: "sequential IDs",
			quests: []Quest{
				{ID: 1},
				{ID: 2},
				{ID: 3},
			},
			want: 4,
		},
		{
			name: "IDs with gap",
			quests: []Quest{
				{ID: 1},
				{ID: 3},
			},
			want: 4,
		},
		{
			name: "unordered IDs",
			quests: []Quest{
				{ID: 5},
				{ID: 2},
				{ID: 9},
			},
			want: 10,
		},
		{
			name: "single quest",
			quests: []Quest{
				{ID: 7},
			},
			want: 8,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NextID(tt.quests)

			if got != tt.want {
				t.Fatalf("NextID() = %d, want %d", got, tt.want)
			}
		})
	}
}
