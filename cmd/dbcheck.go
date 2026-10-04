package cmd

import (
	"fmt"

	"github.com/karryon999/quest/internal/storage"
	"github.com/spf13/cobra"
)

var dbCheckCmd = &cobra.Command{
	Use:   "db-check",
	Short: "Check MySQL connection",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		db, err := storage.OpenMySQL()
		if err != nil {
			return err
		}
		//defer 的意思是： 等当前函数准备结束的时候，再执行这句话。
		defer db.Close()

		fmt.Println("MySQL connection successful.")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(dbCheckCmd)
}
