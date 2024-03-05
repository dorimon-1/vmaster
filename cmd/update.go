/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/dorimon-1/vmaster/config"
	"github.com/dorimon-1/vmaster/utility"
	"github.com/dorimon-1/vmaster/yaml_modifier"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the version of your project",
	Long: `Updates the version of your project. This command will update the version of your project in the helm chart and the version file.
	Usage: vmaster update -e <environment> <serviceName>=<version> <serviceName>=<version> ...
	`,
	Run: func(cmd *cobra.Command, args []string) {
		env, err := cmd.Flags().GetString("environment")
		if err != nil || env == "" {
			fmt.Println("Environment not provided")
			return
		}
		fmt.Println("Environment: ", env)

		arguments := utility.SplitArguments(args)
		fmt.Println(arguments)
		config, err := config.LoadConfig("./config.yaml")
		if err != nil {
			fmt.Println("Error loading config: ", err)
			return
		}

		services, err := yaml_modifier.ParseYAML(config.Enviroments[env].FilePath)
		if err != nil {
			fmt.Println("Error parsing yaml: ", err)
			return
		}

		for key, value := range arguments {
			if _, ok := services[key]; ok {
				services[key].Image.Tag = value
				fmt.Println(services[key].Image.Tag)
			} else {
				fmt.Println("Service not found: ", key)
			}
		}

		if err := yaml_modifier.UpdateYAML(config.Enviroments[env].FilePath, services); err != nil {
			fmt.Println("Error updating yaml: ", err)
			return
		}
		fmt.Println("Done")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	rootCmd.PersistentFlags().
		StringP("environment", "e", "", "The environment to update the version for")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
