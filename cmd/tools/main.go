package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aforamitdev/pdfsync/internal/database"
	"github.com/urfave/cli/v2"
)

type SqlInfo struct {
	FilePath string `json:"file_path"`
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "db",
				Usage: "set sql3 path ",
				Action: func(cCtx *cli.Context) error {
					file_name := cCtx.Args().First()

					if _, err := os.Stat(file_name); err != nil {

					} else if errors.Is(err, os.ErrNotExist) {

						file, err := os.Create("dbconfig.json")

						if err != nil {
							log.Println("error create cinfig ")

						}

						fmt.Println(file)

					} else {
						fmt.Println("file exists ")

						var sqlInfo SqlInfo

						json.Marshal(sqlInfo)
					}

					return nil

				},
			},
			{
				Name:    "migrate",
				Aliases: []string{"t"},
				Usage:   "options for migrtions",
				Subcommands: []*cli.Command{
					{
						Name:  "add",
						Usage: "add new migrations file",
						Action: func(cCtx *cli.Context) error {
							// file_name := cCtx.Args().First()
							// time_nano := time.Now().UnixNano() / 1e6
							database.RunMigrate([]string{"DEV"})

							return nil
						},
					},
					{
						Name:  "remove",
						Usage: "remove an existing template",
						Action: func(cCtx *cli.Context) error {
							fmt.Println("remove migrations  ", cCtx.Args().First())
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
