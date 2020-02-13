package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"rsc.io/quote"
)

func rootCommand() *cobra.Command {
	var rootCmd = &cobra.Command{
		Use:   "go-boilerplate",
		Short: "go-boilerplate is such amaze",
		Long: `go-boilerplate is such amazification that was
   made with love in Go.
   Complete documentation is pending`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hi!")
			color.Cyan(quote.Hello())
		},
	}
	return rootCmd
}

func main() {
	rootCmd := rootCommand()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
