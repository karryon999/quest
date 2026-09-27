package cmd

import (
	"fmt"
	"strconv"
	"time"

	"github.com/karryon999/quest/internal/player"
	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

// 我们之所以在 Cobra 这里使用指针，其中一个重要原因确实是：
// 这个 Command 后面会持续被修改和使用，我们希望大家操作的是同一个 Command 对象。
var doneCmd = &cobra.Command{
	Use:   "done <id>",
	Short: "Complete a quest",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid quest ID %s: must be a number", args[0])
		}
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		found := false
		for i, q := range quests {
			if q.ID == id {
				if q.Completed {
					return fmt.Errorf("quest %d already completed", id)
				}
				//注意：这里不能直接用q.Completed = true
				//因为q是复制了一份当前遍历到的Quest，所以直接用q来修改值的话修改的是副本的值，而不是原对象的值
				quests[i].Completed = true
				now := time.Now()
				quests[i].CompletedAt = &now

				// 将该任务对应的XP加到该玩家已有的XP中，并保存
				p, err := storage.LoadPlayer()
				if err != nil {
					return err
				}
				p.XP += q.XP
				player.UpdateStreak(&p, now)
				if err := storage.SavePlayer(p); err != nil {
					return err
				}

				found = true
				fmt.Printf("Completed quest #%d.\n", q.ID)
				break
			}
		}
		if !found {
			return fmt.Errorf("quest %d not found", id)
		}
		return storage.SaveQuests(quests)
	},
}

func init() {
	rootCmd.AddCommand(doneCmd)
}
