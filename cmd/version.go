package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show current version",
	Long: "show current version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("version called")
		//fmt.Println("system mode:", viper.GetString("system_mode"))

		//fmt.Println("system mode:", rootCmd.Flags().Lookup("system-mode").Value)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	_ = viper.BindPFlag("SYSTEM_MODE", rootCmd.Flags().Lookup("system-mode"))
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// versionCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// versionCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
