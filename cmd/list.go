package cmd

import (
	"fmt"

	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all quests",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		quests, err := storage.LoadQuests()
		if err != nil {
			return err
		}
		if len(quests) == 0 {
			fmt.Println("No quests yet. Add one with:\n\nquest add \"Read paper\"")
			return nil
		}
		fmt.Println("⚔️ Today's Quests")
		for _, q := range quests {
			status := "[ ]"
			if q.Completed {
				status = "[x]"
			}
			boss := ""
			if q.Boss {
				boss = "👹"
			}
			fmt.Printf("%d %s %s +%d XP %s\n", q.ID, status, q.Title, q.XP, boss)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
