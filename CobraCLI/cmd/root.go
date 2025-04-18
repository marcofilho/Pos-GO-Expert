package cmd

import (
	"database/sql"
	"os"

	"github.com/marcofilho/Pos-GO-Expert/CobraCLI/internal/database"
	"github.com/spf13/cobra"

	// sqlite3 driver
	_ "github.com/mattn/go-sqlite3"
)

type RunEFunc func(cmd *cobra.Command, args []string) error

func GetDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
        id TEXT,
        name TEXT,
        description TEXT
    )`)
	if err != nil {
		panic(err)
	}
	return db
}

func GetCategoryDB(db *sql.DB) database.Category {
	return *database.NewCategory(db)
}

var rootCmd = &cobra.Command{
	Use:   "16-CLI",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
