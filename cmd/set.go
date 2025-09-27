package cmd

import (
	"github.com/Chadi-Mangle/hclq/internal/utils"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set <filename> <path> <value>",
	Short: "Set value in a HCL file",
	Long:  "Set value in a HCL file at the specified path.",
	Example: `  hclq set config.hcl resource.instance.name "my-instance"
  hclq set config.hcl resource.instance.count 3`,
	Args: cobra.ExactArgs(3),
	RunE: setCommand,
}

func setCommand(cmd *cobra.Command, args []string) error {
	filename := args[0]
	path := args[1]
	newValue := args[2]

	hclMap, err := utils.ConvertHclFileToMap(filename)
	if err != nil {
		return err
	}

	value, err := utils.FindByPath(hclMap, path)
	if err != nil {
		return err
	}

	value.Set(newValue)

	utils.ConvertMapToHclFile(hclMap, filename)
	return nil
}

func init() {
	rootCmd.AddCommand(setCmd)
}
