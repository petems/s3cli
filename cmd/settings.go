package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// settingsCmd represents the settings command
var settingsCmd = &cobra.Command{
	Use:   "settings",
	Short: "Get a peek at the current settings",
	Long: ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Bucket: %s\n", viper.GetString("bucket"))
		fmt.Printf("Prefix: %s\n", viper.GetString("prefix"))
		fmt.Printf("Region: %s\n", os.Getenv("AWS_REGION"))
	},
}

func init() {
	RootCmd.AddCommand(settingsCmd)
}
