package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/quickbase-textmate/qbtm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var functionsCfg *viper.Viper

var functionsCmd = &cobra.Command{
	Use:   "functions",
	Short: "Build a list of functions for support.function",

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		opts := &qbtm.Opts{}
		err := cliutil.ReadOptions(opts, functionsCfg)
		cliutil.HandleError(cmd, err, "invalid input")

		parser := qbtm.NewFunctionParser()
		err = qbtm.ParseCSV(opts, parser)
		cliutil.HandleError(cmd, err, "error reading csv file")

		b, err := json.Marshal(parser.String())
		cliutil.HandleError(cmd, err, "error encoding functions to json")
		fmt.Println(string(b))
	},
}

func init() {
	var flags *cliutil.Flagger
	functionsCfg, flags = cliutil.AddCommand(rootCmd, functionsCmd, "QUICKBASE_TEXTMATE")
	flags.SetOptions(&qbtm.Opts{})
}
