package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

var (
	hogeCmd = &cobra.Command{
		Use:   "hoge",
		Short: "hoge is wait for 3sec",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunHoge()
		},
	}
)

func RunHoge() error {
	fmt.Println("Start: hoge...")
	defer fmt.Println("End: ...hoge")
	fmt.Println("wait for 3 sec")
	time.Sleep(3 * time.Second)
	return nil
}
