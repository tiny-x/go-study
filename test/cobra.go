package main

import "C"
import (
	"fmt"
	"github.com/howeyc/gopass"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Password: ")
			passwd, err := gopass.GetPasswd()
			fmt.Print(err)
			text := strings.Replace(string(passwd), "\n", "", -1)

			fmt.Println(text)
			return nil
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
