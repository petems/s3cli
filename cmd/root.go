package cmd

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws/session"
)

var cfgFile string

var sess *session.Session

var RootCmd = &cobra.Command{
	Use:   "s3cli",
	Short: "Copy files to and from s3",
	Long: `A simple tool that copies files directly to a preconfigured bucket`,
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(func() { viper.AutomaticEnv(); viper.SetEnvPrefix("s3") })

	sess = session.Must(session.NewSession())
}

func key(filename string) string {
	return path.Join(viper.GetString("prefix"), path.Base(filename))
}
