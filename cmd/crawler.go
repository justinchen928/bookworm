/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package cmd

import (
	"bookworm/src"
	"fmt"

	"github.com/spf13/cobra"
)

// crawlerCmd represents the crawler command
var crawlerCmd = &cobra.Command{
	Use:   "crawler",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("crawler called")
		novel := src.Novel{}
		crawler := src.NewBookCrawler(args[0])
		crawler.SetCollector(src.NewCollector())
		crawler.Crawl(&novel)
		novel.SaveToTxt(args[1])
	},
}

func init() {
	rootCmd.AddCommand(crawlerCmd)
	crawlerCmd.Flags().BoolP("path", "p", false, "specify txt path")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// crawlerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// crawlerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
