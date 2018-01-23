package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	// ServicesCmd represents the base command when called without any subcommands
	ServicesCmd = &cobra.Command{
		Use:   "services",
		Short: "",
		Long:  "",
	}
	ConfPath string
)

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := ServicesCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	ServicesCmd.PersistentFlags().StringVarP(&ConfPath, "conf", "c", "", "path to config")
}
