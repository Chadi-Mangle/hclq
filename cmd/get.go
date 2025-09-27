package cmd

import (
	"fmt"

	"github.com/Chadi-Mangle/hclq/internal/utils"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get <filename> [path]",
	Short: "Get value from a HCL file",
	Long:  "Get value from a HCL file. If no path is specified, returns the entire file content.",
	Example: `  hclq get config.hcl
  hclq get config.hcl resource.instance.name`,
	Args: cobra.RangeArgs(1, 2),
	RunE: getCommand,
}

func getCommand(cmd *cobra.Command, args []string) error {
	filename := args[0]

	hclMap, err := utils.ConvertHclFileToMap(filename)
	if err != nil {
		return err
	}

	if len(args) == 1 {
		hclBytes, err := utils.ConvertMapToHcl(hclMap)
		if err != nil {
			return err
		}
		fmt.Println(string(hclBytes))
		return nil
	}

	path := args[1]
	value, err := utils.FindByPath(hclMap, path)
	if err != nil {
		return err
	}

	valueBytes, err := value.GetHcl()
	if err != nil {
		return err
	}

	fmt.Println(string(valueBytes))
	return nil
}

func init() {
	rootCmd.AddCommand(getCmd)
}
