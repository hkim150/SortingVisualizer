package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("size", "s", 30, "size of the array to sort")
	rootCmd.Flags().StringP("sorter", "a", "bubble", "sorting algorithm - choose from [bubble]")
}

var rootCmd = &cobra.Command{
	Short: "Visualize Sorting",
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		algorithm, _ := cmd.Flags().GetString("sorter")

		fmt.Println(size, algorithm)
		// err := sorter.Sort(data)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
