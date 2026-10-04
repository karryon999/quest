package cmd

import (
	"fmt"

	"github.com/karryon999/quest/internal/quest"
	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <title>",
	Short: "Add a new quest",
	Args:  cobra.ExactArgs(1), // 限制参数数量必须刚好是 1 个
	//json存储版本：
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	q := quest.New(args[0], addXP, addBoss)
	//	// quests指的是加载的已有的任务，q是要新增的任务
	//	quests, err := storage.LoadQuests()
	//	if err != nil {
	//		return err
	//	}
	//	// q的ID是根据已有的任务的ID算出来的，所以这里传入的是quests
	//	q.ID = quest.NextID(quests)
	//	quests = append(quests, q)
	//	err = storage.SaveQuests(quests)
	//	if err != nil {
	//		return err
	//	}
	//	fmt.Printf("Added quest #%d:%s\n", q.ID, q.Title)
	//	return nil
	//},

	//mysql存储版本：
	RunE: func(cmd *cobra.Command, args []string) error {
		q := quest.New(args[0], addXP, addBoss)

		db, err := storage.OpenMySQL()
		if err != nil {
			return err
		}
		defer db.Close()

		id, err := storage.AddQuestMySQL(db, q)
		if err != nil {
			return err
		}

		fmt.Printf("Added quest #%d: %s\n", id, q.Title)
		return nil
	},
}

var addXP int
var addBoss bool

// 这个 package 被加载时，Go 会自动执行 init()函数
func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().IntVar(&addXP, "xp", 30, "XP reward for the quest")
	//没有写 --boss -> 使用默认值 false
	addCmd.Flags().BoolVar(&addBoss, "boss", false, "Add boss for the quest")
}
