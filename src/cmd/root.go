package cmd

import (
	"github.com/spf13/cobra"
	"picket-main-service/src/config"
)

type funcCmd = func(config config.IConfig) *cobra.Command

func GetRoot(config config.IConfig) *cobra.Command {
	cmd := []funcCmd{migrateUp, migrateDown, migrateRefresh, server}
	root := &cobra.Command{}

	for _, item := range cmd {
		root.AddCommand(item(config))
	}

	return root
}
