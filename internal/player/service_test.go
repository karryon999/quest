package player

import (
	"testing"
	"time"
)

func timePtr(t time.Time) *time.Time {
	return &t
}

func TestUpdateStreak(t *testing.T) {
	tests := []struct {
		name        string
		player      Player
		now         time.Time
		wantCurrent int
		wantBest    int
	}{
		{
			name:   "first completion",
			player: Player{},
			now: time.Date(
				2026, 9, 27,
				10, 0, 0, 0,
				time.UTC,
			),
			wantCurrent: 1,
			wantBest:    1,
		},
		{
			name: "same day does not increase streak",
			player: Player{
				CurrentStreak: 1,
				BestStreak:    1,
				LastCompletedDate: timePtr(time.Date(
					2026, 9, 27,
					8, 0, 0, 0,
					time.UTC,
				)),
			},
			now: time.Date(
				2026, 9, 27,
				20, 0, 0, 0,
				time.UTC,
			),
			wantCurrent: 1,
			wantBest:    1,
		},
		{
			name: "next day increases streak",
			player: Player{
				CurrentStreak: 1,
				BestStreak:    1,
				LastCompletedDate: timePtr(time.Date(
					2026, 9, 26,
					10, 0, 0, 0,
					time.UTC,
				)),
			},
			now: time.Date(
				2026, 9, 27,
				10, 0, 0, 0,
				time.UTC,
			),
			wantCurrent: 2,
			wantBest:    2,
		},
		{
			name: "missed day resets current streak",
			player: Player{
				CurrentStreak: 5,
				BestStreak:    5,
				LastCompletedDate: timePtr(time.Date(
					2026, 9, 25,
					10, 0, 0, 0,
					time.UTC,
				)),
			},
			now: time.Date(
				2026, 9, 27,
				10, 0, 0, 0,
				time.UTC,
			),
			wantCurrent: 1,
			wantBest:    5,
		},
		{
			name: "new streak updates best streak",
			player: Player{
				CurrentStreak: 3,
				BestStreak:    3,
				LastCompletedDate: timePtr(time.Date(
					2026, 9, 26,
					10, 0, 0, 0,
					time.UTC,
				)),
			},
			now: time.Date(
				2026, 9, 27,
				10, 0, 0, 0,
				time.UTC,
			),
			wantCurrent: 4,
			wantBest:    4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := tt.player

			UpdateStreak(&p, tt.now)

			if p.CurrentStreak != tt.wantCurrent {
				t.Errorf(
					"CurrentStreak = %d, want %d",
					p.CurrentStreak,
					tt.wantCurrent,
				)
			}

			if p.BestStreak != tt.wantBest {
				t.Errorf(
					"BestStreak = %d, want %d",
					p.BestStreak,
					tt.wantBest,
				)
			}

			if p.LastCompletedDate == nil {
				t.Fatal("LastCompletedDate is nil")
			}

			if !p.LastCompletedDate.Equal(tt.now) {
				t.Errorf(
					"LastCompletedDate = %v, want %v",
					*p.LastCompletedDate,
					tt.now,
				)
			}
		})
	}
}
