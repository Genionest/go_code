package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "root",
	Short: "short desc",
	Long:  "long desc",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root cmd run begin")
		fmt.Println(
			cmd.PersistentFlags().Lookup("viper").Value,
			cmd.PersistentFlags().Lookup("config").Value,
			cmd.PersistentFlags().Lookup("author").Value,
			cmd.Flags().Lookup("source").Value,
		)
		fmt.Println("root cmd run end")
	},
}

func Execute() {
	rootCmd.Execute()
}

var author string
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().Bool("viper", true, "usage")
	// rootCmd.PersistentFlags().BoolP("viper", "v", true, "usage")
	// rootCmd.PersistentFlags().StringP("author", "a", "Your Name", "usage")
	rootCmd.PersistentFlags().StringVarP(&author, "author", "a", "Your Name", "usage")
	rootCmd.Flags().StringP("source", "s", "", "usage")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cobra")
	}
	// 检查环境变量，将配置的键值添加到viper
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println(err)
	}
	fmt.Println("usering config file:", viper.ConfigFileUsed())
}

func main() {
	Execute()
}
