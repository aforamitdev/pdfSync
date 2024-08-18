package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/httpfs"

	_ "embed"

	// _ "github.com/mattn/go-sqlite3"
	"github.com/urfave/cli/v2"
)

type SqlInfo struct {
	FilePath string `json:"file_path"`
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{

			{
				Name:    "migrate",
				Aliases: []string{"t"},
				Usage:   "options for migrtions",
				Subcommands: []*cli.Command{
					{
						Name:  "commit",
						Usage: "commit migrations to sqllite",
						Action: func(cCtx *cli.Context) error {

							db, err := sql.Open("sqlite3", "file:app.db")
							if err != nil {
								fmt.Println(err)
							}

							migrations, err := httpfs.New(http.Dir("./"), "migrations")
							if err != nil {
								fmt.Println(err)
							}
							defer db.Close()

							instance, err := sqlite3.WithInstance(db, new(sqlite3.Config))

							if err != nil {
								fmt.Println(err)
							}

							m, err := migrate.NewWithInstance("httpfs", migrations, "sqlite3", instance)

							if err != nil {
								fmt.Println("error")
								fmt.Println(err)
							}
							// fmt.Println(m.)
							if err := m.Up(); err != nil {
								fmt.Println("error")
								fmt.Println(err)
							}
							return nil
						},
					},
					{
						Name:  "create",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							migrations_name := cCtx.Args().First()
							if migrations_name == "" {
								return cli.Exit("migration name is required", 1)
							}
							timesec := time.Now().Local().Unix()
							up_file := fmt.Sprintf("%d_%s.up.sql", timesec, migrations_name)
							down_file := fmt.Sprintf("%d_%s.down.sql", timesec, migrations_name)
							migratSli := []string{up_file, down_file}

							for _, file := range migratSli {
								fs, err := os.OpenFile(fmt.Sprintf("migrations/%s", file), os.O_RDONLY|os.O_CREATE, 0666)
								if err != nil {
									fmt.Println(err)
								}
								defer fs.Close()

							}

							return nil
						},
					},
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
