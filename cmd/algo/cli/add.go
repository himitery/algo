package cli

import (
	"algo/cmd/algo/app"
	"algo/internal/usecase"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new algorithm problem",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		lang, _ := cmd.Flags().GetString("lang")
		target, _ := cmd.Flags().GetString("platform")
		id := args[0]

		app.New(fx.Invoke(func(algoUsecase usecase.Algo) {
			algoUsecase.Add(lang, target, id)
		}))
	},
}

func init() {
	addCmd.Flags().StringP("lang", "l", "python", "programming language")
	addCmd.Flags().StringP("platform", "p", "baekjoon", "platform")
}
