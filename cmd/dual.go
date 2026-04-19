package cmd

import (
	"errors"
	"fmt"

	"github.com/megaproaktiv/ltail/bubble"
	"github.com/megaproaktiv/ltail/config"
	"github.com/spf13/cobra"
)

var dualConfig1 config.Configuration
var dualConfig2 config.Configuration
var dualOutputConfig1 config.OutputConfiguration
var dualOutputConfig2 config.OutputConfiguration

var dualCommand = &cobra.Command{
	Use:   "dual <log-group-1> <log-group-2>",
	Short: "Watch two log groups in split view",
	Long:  "Watch two log groups side by side in a fullscreen split terminal view using Bubble Tea TUI",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 2 {
			return errors.New("dual command requires two log group arguments")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		dualConfig1.Group = args[0]
		dualConfig2.Group = args[1]

		err := bubble.Run(
			args[0],
			args[1],
			&awsConfig,
			&dualConfig1,
			&dualConfig2,
			&dualOutputConfig1,
			&dualOutputConfig2,
		)

		if err != nil {
			fmt.Printf("Error running dual view: %v\n", err)
		}
	},
}

func init() {
	// First pane flags
	dualCommand.Flags().StringVar(&dualConfig1.Prefix, "prefix1", "", "log stream prefix filter for first pane")
	dualCommand.Flags().StringVar(&dualConfig1.Filter, "filter1", "", "event filter pattern for first pane")
	dualCommand.Flags().StringVar(&dualConfig1.Start, "start1", "", "start time for first pane")
	dualCommand.Flags().BoolVar(&dualOutputConfig1.Shorten, "shorten1", false, "shorten lines for first pane")

	// Second pane flags
	dualCommand.Flags().StringVar(&dualConfig2.Prefix, "prefix2", "", "log stream prefix filter for second pane")
	dualCommand.Flags().StringVar(&dualConfig2.Filter, "filter2", "", "event filter pattern for second pane")
	dualCommand.Flags().StringVar(&dualConfig2.Start, "start2", "", "start time for second pane")
	dualCommand.Flags().BoolVar(&dualOutputConfig2.Shorten, "shorten2", false, "shorten lines for second pane")

	// Common flags (apply to both)
	dualCommand.Flags().BoolVarP(&dualOutputConfig1.Shorten, "shorten", "s", false, "shorten lines for both panes")
	dualCommand.Flags().BoolVarP(&dualOutputConfig2.Shorten, "shorten-both", "", false, "alias for --shorten")
}