package cmd

import (
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/crypto/ssh/terminal"
	"golang.org/x/sys/unix"
)

var rootCmd = &cobra.Command{
	Use:   "jiritsu",
	Short: "jiritsu web service",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if !terminal.IsTerminal(unix.Stdout) {
			logrus.SetFormatter(&logrus.JSONFormatter{})
		} else {
			logrus.SetFormatter(&logrus.TextFormatter{
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339Nano,
			})
		}

		if verbose, _ := cmd.Flags().GetBool("verbose"); verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
	},
}
