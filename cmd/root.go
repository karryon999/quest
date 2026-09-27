// cmd负责CLI输入
// 我们其实是在做命令行的命令，就是使用quest xxx就可以在终端执行命令，类似git xx的命令
// 然后我们会有quest add xx，quest list， quest delete xx等CLI命令
// 现在我们在开发阶段不能直接用quest命令，而是要用go run .来代替quest

package cmd

import "github.com/spf13/cobra"

// rootCmd是一个指向cobra.Command{}结构体的指针，rootCmd 的类型是 *cobra.Command
var rootCmd = &cobra.Command{
	Use:   "quest",
	Short: "Turn your boring TODO list into an RPG adventure.",
}

func Execute() error {
	return rootCmd.Execute()
}
