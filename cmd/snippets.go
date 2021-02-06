package cmd

import (
	"fmt"

	"github.com/cpliakas/cliutil"
	"github.com/cpliakas/quickbase-textmate/qbtm"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var snippetsCfg *viper.Viper

var snippetsCmd = &cobra.Command{
	Use:   "snippets",
	Short: "Build the snippets.json file",

	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},

	Run: func(cmd *cobra.Command, args []string) {
		opts := &qbtm.Opts{}
		err := cliutil.ReadOptions(opts, snippetsCfg)
		cliutil.HandleError(cmd, err, "invalid input")

		parser := qbtm.NewSnippetParser()
		err = qbtm.ParseCSV(opts, parser)
		cliutil.HandleError(cmd, err, "error reading csv file")
		fmt.Println(parser)
	},
}

func init() {
	var flags *cliutil.Flagger
	snippetsCfg, flags = cliutil.AddCommand(rootCmd, snippetsCmd, "QUICKBASE_TEXTMATE")
	flags.SetOptions(&qbtm.Opts{})
}
