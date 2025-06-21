package cmd

import (
	"fmt"
	"os"
	"sortingvisualizer/internal"
	"sortingvisualizer/internal/array"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.Flags().IntP("size", "s", 30, "size of the array to sort")
	rootCmd.Flags().StringP("algorithm", "a", "bubble", fmt.Sprintf("sorting algorithm - choose from [%v]", internal.Algorithms()))
}

var rootCmd = &cobra.Command{
	Short: "Visualize Sorting",
	Run: func(cmd *cobra.Command, args []string) {
		size, _ := cmd.Flags().GetInt("size")
		algo, _ := cmd.Flags().GetString("algorithm")

		if !(1 <= size && size <= 100) {
			fmt.Println("Size must be between 1 to 100")
			return
		}

		arr := array.NewArray(size)
		arr.Print()

		err := internal.Sort(arr, algo)
		if err != nil {
			fmt.Println("Error sorting the data: ", err)
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
