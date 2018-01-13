package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var restoreCmd = &cobra.Command{
	Use:   "restore [filename]",
	Short: "",
	Long: ``,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		downloader := s3manager.NewDownloader(sess)

		file, err := os.Create(filename)
		if err != nil {
		    return err
		}

		n, err := downloader.Download(file, &s3.GetObjectInput{
		    Bucket: aws.String(viper.GetString("bucket")),
		    Key:    aws.String(key(filename)),
		})
		if err == nil {
			fmt.Printf("file downloaded, %d bytes\n", n)
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(restoreCmd)
}
