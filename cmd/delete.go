package cmd

import (
	"fmt"
	"strconv"

	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a quest",
	Args:  cobra.ExactArgs(1),
	//json存储版本：
	//RunE: func(cmd *cobra.Command, args []string) error {
	//	// 把字符串转换成整数
	//	id, err := strconv.Atoi(args[0])
	//	if err != nil {
	//		return fmt.Errorf("invalid quest ID %s: must be a number", args[0])
	//	}
	//	quests, err := storage.LoadQuests()
	//	if err != nil {
	//		return err
	//	}
	//	found := false
	//	for i, q := range quests {
	//		if q.ID == id {
	//			// slice没有delete操作，所以采用前后拼接法
	//			//quests[i+1:] 本身是一整个 slice。... 的意思可以暂时理解为：
	//			//把这个 slice 拆开，把里面的元素一个一个追加进去。
	//			quests = append(quests[:i], quests[i+1:]...)
	//			fmt.Printf("Deleted quest #%d.\n", q.ID)
	//			found = true
	//			break
	//		}
	//	}
	//	if !found {
	//		return fmt.Errorf("quest %d not found", id)
	//	}
	//	if err := storage.SaveQuests(quests); err != nil {
	//		return err
	//	}
	//	return nil
	//},

	//mysql存储版本：
	RunE: func(cmd *cobra.Command, args []string) error {
		//先将用户输入的要删除的任务id转为int
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid quest ID %s: must be a number", args[0])
		}
		//连接数据库
		db, err := storage.OpenMySQL()
		if err != nil {
			return err
		}
		defer db.Close()

		deleted, err := storage.DeleteQuestMySQL(db, id)
		if err != nil {
			return err
		}
		if !deleted {
			return fmt.Errorf("quest %d not found", id)
		}
		fmt.Printf("Deleted quest #%d.\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
