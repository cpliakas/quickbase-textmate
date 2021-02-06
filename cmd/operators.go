package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/quickbase-textmate/qbtm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var operatorsCfg *viper.Viper

var operatorsCmd = &cobra.Command{
	Use:   "operators",
	Short: "Build a list of operators for support.function",

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		opts := &qbtm.Opts{}
		err := cliutil.ReadOptions(opts, operatorsCfg)
		cliutil.HandleError(cmd, err, "invalid input")

		parser := qbtm.NewOperationParser()
		err = qbtm.ParseCSV(opts, parser)
		cliutil.HandleError(cmd, err, "error reading csv file")

		b, err := json.Marshal(parser.String())
		cliutil.HandleError(cmd, err, "error encoding operators to json")
		fmt.Println(string(b))
	},
}

func init() {
	var flags *cliutil.Flagger
	operatorsCfg, flags = cliutil.AddCommand(rootCmd, operatorsCmd, "QUICKBASE_TEXTMATE")
	flags.SetOptions(&qbtm.Opts{})
}
