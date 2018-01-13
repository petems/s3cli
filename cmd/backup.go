package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var backupCmd = &cobra.Command{
	Use:   "backup [filename]",
	Short: "",
	Long: ``,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		filename := args[0]

		uploader := s3manager.NewUploader(sess)
		file, err  := os.Open(filename)
		if err != nil {
			return err
		}

		d, err := uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(viper.GetString("bucket")),
			Key:    aws.String(key(filename)),
			Body:   file,
		})
		if err == nil {
			fmt.Printf("File uploaded to: %s (version: %s)\n", d.Location, d.VersionID)			
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(backupCmd)
}
