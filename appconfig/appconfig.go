package appconfig

import (
	"fmt"

	"github.com/spf13/cobra"
)

func cmdAppConfig() *cobra.Command {
	// o := &Options{}
	cmd := &cobra.Command{
		Use:   "appconfig",
		Short: "aws appconfig",
		Long:  "aws appconfig command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Say hello. This is my first go library")
		},
	}

	// 	// o.BaseOptions.AddBaseFlags(cmd)
	return cmd
}
