package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	fugaCmd = &cobra.Command{
		Use:   "fuga",
		Short: "fuga is wait for 3sec",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunFuga()
		},
	}
)

func RunFuga() error {
	fmt.Println("Start: fuga...")
	defer fmt.Println("End: ...fuga")
	fmt.Println("wait for 5 sec")
	time.Sleep(5 * time.Second)
	return nil
}
