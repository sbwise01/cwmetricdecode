/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
	"github.com/spf13/cobra"
	protomessage "github.com/open-telemetry/opentelemetry-proto/gen/go/collector/metrics/v1"
)

// decodeCmd represents the decode command
var decodeCmd = &cobra.Command{
	Use:   "decode",
	Short: "Command to decode Cloudwatch metrics stream data written to S3 buckets",
	Long: `This is a golang command application that will
take an argument for a file that was downloaded from an S3
oject created in an S3 bucket by Cloudwatch metrics streams
and use protobuf to decode the one to many Metrics messages it contains.`,
	Args: cobra.ExactArgs(1),
	Example: `cwmetricdecode decode <filename>
cwmetricdecode decode ./test/data/MetricStreams-QuickFull-k2WzmU-ZRJgfgu7`,
	Run: func(cmd *cobra.Command, args []string) {
		var metricsFilename string

		metricsFilename = args[0]

		contents, err := ioutil.ReadFile(metricsFilename)
		if err != nil {
			fmt.Printf("Error reading file contents:  %v\n", err)
		}
		fmt.Printf("Content length of file is %d\n",len(contents))

		err = deserialize(contents)
		if err != nil {
			fmt.Printf("Error decoding file contents:  %v\n", err)
		}		

	},
}

func init() {
	rootCmd.AddCommand(decodeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// decodeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// decodeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func deserialize(content []byte) error {
	var metrics protomessage.ExportMetricsServiceRequest

	finishedDecoding := false

	buffer := proto.NewBuffer(content)
	for finishedDecoding != true {
		err := buffer.DecodeMessage(&metrics)
		if err != nil {
			finishedDecoding = true
		} else {
			fmt.Printf("Content: %v\n", metrics.String())
			metrics.Reset()
		}
	}

	return nil
}
