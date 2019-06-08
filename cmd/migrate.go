package cmd

import (
	"github.com/eltrufas/jiritsu/db"
	"github.com/eltrufas/jiritsu/migrations"
	"github.com/spf13/cobra"
	errors "golang.org/x/xerrors"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	RunE: func(cmd *cobra.Command, args []string) error {
		dbCfg, err := db.InitConfig()
		if err != nil {
			return errors.Errorf("Unable to load db config: %w", err)
		}

		db, err := db.New(dbCfg)
		if err != nil {
			return errors.Errorf("Unable to open db: %w", err)
		}
		return migrations.Migrate(db.DB)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
