package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/guumaster/surveygen/pkg/generator"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "surveygen",
	Short: "Generate go code for surveys",
	Long: `
surveygen is a CLI tool to generate go code for surveys from a yaml definition file.
`,
	SilenceUsage: true,
	RunE: func(cmd *cobra.Command, args []string) error {
		paths, _ := cmd.Flags().GetStringSlice("path")
		if len(paths) == 0 {
			return errors.New("path flag is mandatory")
		}
		g := generator.New()
		for _, p := range paths {
			err := g.Generate(p)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {

	defaultPath, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	rootCmd.PersistentFlags().StringSliceP("path", "p", []string{defaultPath}, "path to look for survey definitions")
}
