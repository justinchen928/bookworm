/*
Copyright © 2022 rfaychen justin928501@gmail.com

*/
package commands

import (
	"bookworm/converter"
	"bookworm/scraper"
	"errors"
	"log"

	"github.com/spf13/cobra"
)

func startScrap(cmd *cobra.Command, args []string) {

	log.Println("scraper called")

	is_to_epub, _ := cmd.Flags().GetBool("epub")
	is_to_kepub, _ := cmd.Flags().GetBool("kepub")
	is_to_json, _ := cmd.Flags().GetBool("json")
	is_o, _ := cmd.Flags().GetBool("output")

	scraper_config := scraper.Config{
		URL: args[0],
	}

	scraper := scraper.New(scraper_config)
	scraper.Start()

	output_dir := "./"
	if is_o && len(args) > 1 {
		output_dir = args[1]
	}
	converter_config := converter.Config{
		OutputDir: output_dir,
	}
	converter := converter.New(converter_config)

	if is_to_kepub {
		converter.ConvertAndSave(scraper.GetNovel(), "kepub")
	}

	if is_to_epub {
		converter.ConvertAndSave(scraper.GetNovel(), "epub")
	}

	if is_to_json {
		converter.ConvertAndSave(scraper.GetNovel(), "json")
	}
}

func (b *commandsBuilder) newScrapCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "scraper [link]",
		Short: "Scrape web novels of certain websites and save to file",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires target link\n")
			}
			return nil
		},
		Run: startScrap,
	}

	cmd.Flags().Bool("epub", false, "save the novel as epub")
	cmd.Flags().Bool("kepub", false, "save the novel as kepub (kobo epub)")
	cmd.Flags().Bool("json", false, "save the novel as JSON")
	cmd.Flags().BoolP("output", "o", false, "output path")

	return cmd
}
