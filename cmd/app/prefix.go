/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package app

import (
	"fmt"
	"sync"

	"github.com/spf13/cobra"

	"github.com/dorimon-1/vmaster/pkg/config"
	"github.com/dorimon-1/vmaster/pkg/utility"
	"github.com/dorimon-1/vmaster/pkg/yaml_modifier"
)

// prefixCmd represents the prefix command
var prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Updates your environments with the given prefix",
	Long: `Updates your environments with the given prefix. 
	This command will update the version of your project in the helm chart and the version file.`,

	Run: func(cmd *cobra.Command, args []string) {
		prefix, err := cmd.Flags().GetString("prefix")
		if err != nil || prefix == "" {
			fmt.Println("Prefix not provided")
			return
		}

		config, err := config.LoadConfig("./config.yaml")
		if err != nil {
			fmt.Println("Error loading config: ", err)
			return
		}
		arguments := utility.SplitArguments(args)

		if err := doPrefix(prefix, config, arguments); err != nil {
			fmt.Println("Error updating prefix: ", err)
			return
		}
	},
}

func doPrefix(prefix string, c *config.Config, arguments map[string]string) error {
	var wg sync.WaitGroup
	fmt.Println("Environments: ", c.Environments)
	for _, v := range c.Environments {
		if v.VersionPrefix == prefix {
			wg.Add(1)
			go func(env *config.Environment) {
				defer wg.Done()
				if err := performUpdate(arguments, env); err != nil {
					fmt.Println("Error updating prefix: ", err)
					return
				}
			}(&v)
		}
	}
	wg.Wait()

	return nil
}

func performUpdate(arguments map[string]string, env *config.Environment) error {
	fmt.Println("Updating: ", env.FilePath)
	service, err := yaml_modifier.ParseYAML(env.FilePath)
	if err != nil {
		return fmt.Errorf("Error parsing yaml: %v", err)
	}
	for serviceName, serviceVersion := range arguments {
		if _, ok := service[serviceName]; ok {
			if service[serviceName].Image.Tag == serviceName {
				continue
			}
			service[serviceName].Image.Tag = serviceVersion
		} else {
			fmt.Println("Service not found: ", serviceName)
			continue
		}
	}
	if err := yaml_modifier.UpdateYAML(env.FilePath, service); err != nil {
		return fmt.Errorf("Error updating yaml: %v", err)
	}
	fmt.Println("Updated: ", env.FilePath)
	return nil
}

func init() {
	updateCmd.AddCommand(prefixCmd)
	updateCmd.PersistentFlags().StringP("prefix", "p", "", "Prefix for the version")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// prefixCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// prefixCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
