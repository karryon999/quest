//stats是statistics的常见缩写，表示统计数据/属性数据，在游戏里很常见
//quest stats表示查看玩家的统计信息

package cmd

import (
	"fmt"

	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

var statsCmd = &cobra.Command{
	Use:   "stats",
	Short: "Show adventurer stats", //显示冒险者属性
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, err := storage.LoadPlayer()
		if err != nil {
			return err
		}
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		completedCount := 0
		for _, q := range quests {
			if q.Completed {
				completedCount++
			}
		}
		fmt.Println("📊 Adventurer Stats")
		fmt.Printf("Level: %d\n", p.Level())
		fmt.Printf("XP: %d\n", p.XP)
		fmt.Printf("Completed Quests: %d\n", completedCount)
		fmt.Printf("Current Streak: %d\n", p.CurrentStreak)
		fmt.Printf("Best Streak: %d\n", p.BestStreak)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(statsCmd)
}
