package cmd

import (
	"github.com/marcofilho/Pos-GO-Expert/CobraCLI/internal/database"
	"github.com/spf13/cobra"
)

func newListCmd(categoryDb database.Category) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List categories",
		Long:  `List categories`,
		RunE:  runGetCategories(categoryDb),
	}
}

func runGetCategories(categoryDb database.Category) RunEFunc {
	return func(cmd *cobra.Command, args []string) error {
		_, err := categoryDb.FindAll()
		if err != nil {
			return err
		}
		return nil
	}
}

func init() {
	listCmd := newListCmd(GetCategoryDB(GetDb()))
	categoryCmd.AddCommand(listCmd)
}
