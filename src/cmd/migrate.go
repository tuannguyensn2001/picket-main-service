package cmd

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/spf13/cobra"
	"log"
	"os"
	"picket-main-service/src/config"
)

func setupM(config config.IConfig) *migrate.Migrate {
	db, err := config.GetDB().DB()
	if err != nil {
		log.Fatalln("err connect to db")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	dir, _ := os.Getwd()
	path := "file://" + dir + "/src/database/migrations"

	m, err := migrate.NewWithDatabaseInstance(path, "pgsql", driver)

	if err != nil {
		fmt.Println("err migrate up 1 ", err)
	}

	return m
}

func migrateUp(config config.IConfig) *cobra.Command {
	return &cobra.Command{

		Use: "migrate-up",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateDown(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-down",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Steps(-1)

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}

func migrateRefresh(config config.IConfig) *cobra.Command {
	return &cobra.Command{
		Use: "migrate-refresh",
		Run: func(cmd *cobra.Command, args []string) {

			m := setupM(config)

			err := m.Down()
			if err != nil {
				log.Fatalln(err)
			}
			err = m.Up()

			if err != nil {
				fmt.Println("err migrate up 2", err)
			}
		},
	}
}
