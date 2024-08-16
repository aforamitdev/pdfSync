package database

import (
	"database/sql"
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DEV  = "dev"
	PROD = "prod"
	TEDT = "test"
)

var (
	config_devloment  map[string]string
	config_production map[string]string
)

func RunMigrate(args []string) {

	// args = args[2:]

	switch args[0] {
	case DEV:
		migrate_DB(config_devloment)
	case PROD:
		migrate_DB(config_production)
	default:
		migrate_DB(config_devloment)

	}

}

func migrate_DB(config map[string]string) {

	var migrations []string
	var completed []string

	files, err := filepath.Glob("./migrations/*.sql")

	if err != nil {
		log.Printf("error running restore %s", err)
		return
	}
	sort.Strings(files)

	db, err := OpenSqlDb()

	if err != nil {
		fmt.Println("")
	}
	log.Printf("%d\n", db.Stats().Idle)

	if err != nil {
		log.Println("no database found")
	} else {
		migrations = readMetadata(db)
	}

	fmt.Println(files, "files")
	for _, file := range files {
		file_name := filepath.Base(file)
		if !contains(file_name, migrations) {
			log.Printf("running migrations %s", file_name)
			result, err := runCommand("app.db", file_name)
			if err != nil || strings.Contains(string(result), "ERROR") {
				if err == nil {
					err = fmt.Errorf("\n%s", string(result))
				}
				completed = append(completed, file_name)
				// If at any point we fail, log it and break
				log.Printf("ERROR loading sql migration:%s\n", err)
				log.Printf("All further migrations cancelled\n\n")
				break
			}

		}
	}
	if len(completed) > 0 {
		writeMetadata(config, completed, db)

	}

}

func readMetadata(db *sql.DB) []string {

	var migrations []string

	sql := "select migration_version from sync_metadata order by id desc;"

	rows, err := db.Query(sql)
	if err != nil {
		log.Printf("dabase error %s", err)
		return migrations
	}
	defer rows.Close()
	for rows.Next() {
		var migration string
		err := rows.Scan(&migration)
		if err != nil {
			log.Printf("database error %s", err)
			return migrations
		}
		migrations = append(migrations, migration)
	}
	return migrations
}

func writeMetadata(config map[string]string, migrations []string, db *sql.DB) {

	for _, m := range migrations {
		sql := "insert into sync_metadata(updated_at,sync_version,migration_version,status) VALUES(NOW(),$1,$2,100);"
		result, err := db.Exec(sql, 1.0, m)
		if err != nil {
			log.Printf("database error %s %s", err, result)
		}
	}

}

func contains(s string, a []string) bool {
	for _, k := range a {
		if s == k {
			return true
		}
	}
	return false
}

func runCommand(db string, file_name string) ([]byte, error) {

	fmt.Println(db, file_name, fmt.Sprintf("./migrations/%s", file_name))
	cmd := exec.Command("sqlite3", db, "<", fmt.Sprintf("./migrations/%s", file_name))
	output, err := cmd.CombinedOutput()
	fmt.Println(string(output))
	if err != nil {
		return output, err
	}
	return output, nil
}
