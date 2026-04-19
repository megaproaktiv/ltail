package cmd

import (
	"github.com/megaproaktiv/ltail/config"
	"github.com/spf13/cobra"
)

// LtailCommand is the main top-level command
var LtailCommand = &cobra.Command{
	Use:   "ltail <command>",
	Short: "A fast, multipurpose tool for AWS CloudWatch Logs",
	Long:  "ltail is a fast, multipurpose tool for AWS CloudWatch Logs.",
	Example: `  ltail groups
  ltail streams production
  ltail watch production`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.HelpFunc()(cmd, args)
	},
}

var awsConfig config.AWSConfiguration

func init() {
	LtailCommand.AddCommand(groupsCommand)
	LtailCommand.AddCommand(streamsCommand)
	LtailCommand.AddCommand(versionCommand)
	LtailCommand.AddCommand(watchCommand)
	LtailCommand.AddCommand(getCommand)
	LtailCommand.AddCommand(dualCommand)
	LtailCommand.PersistentFlags().StringVar(&awsConfig.Endpoint, "endpoint-url", "", "override default endpoint URL")
	LtailCommand.PersistentFlags().StringVar(&awsConfig.Region, "region", "", "override profile AWS region")
	LtailCommand.PersistentFlags().StringVar(&awsConfig.Profile, "profile", "", "override default AWS profile")
}
