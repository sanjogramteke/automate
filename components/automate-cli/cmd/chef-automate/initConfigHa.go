// Copyright © 2017 Chef Software

package main

import (
	"errors"

	"github.com/chef/automate/components/automate-cli/pkg/docs"
	pgc "github.com/chef/automate/components/automate-cli/pkg/pullandgenerateconfig"
	"github.com/chef/automate/components/automate-cli/pkg/status"
	"github.com/spf13/cobra"
)

func init() {
	initConfigHACmd.PersistentFlags().StringVar(
		&pgc.InitConfigHAPathFlags.Path,
		"file",
		"config.toml",
		"File path to write the config")
	initConfigHACmd.SetUsageTemplate(UsageTemplate)

	initConfigHACmd.PersistentFlags().StringVar(
		&pgc.InitConfigHabA2HAPathFlag.A2haDirPath,
		"path",
		"/hab/a2_deploy_workspace/",
		"a2ha hab workspace dir path")
	initConfigHACmd.SetUsageTemplate(UsageTemplate)
	initConfigHACmd.PersistentFlags().SetAnnotation("path", docs.Compatibility, []string{docs.CompatiblewithHA})
	RootCmd.AddCommand(initConfigHACmd)
}

var initConfigHACmd = &cobra.Command{
	Use:   "init-config-ha",
	Short: "Initialize default config for Automate HA",
	Long:  "Initialized default configuration for HA and save it to a file.",
	Annotations: map[string]string{
		NoCheckVersionAnnotation: NoCheckVersionAnnotation,
		docs.Compatibility:       docs.CompatiblewithHA,
	},
	RunE: runInitConfigHACmd,
}

func runInitConfigHACmd(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		msg := "one argument expected, please refer help doc."
		writer.Printf("%s\n\n%s\n", msg, UsageTemplate)
		return nil
	} else if args[0] == "aws" {
		writer.Printf("Generating initial automate high availability configuration for AWS deployment\n")
		return runInitConfigAwsHACmd()
	} else if args[0] == "existing_infra" {
		writer.Printf("Generating initial automate high availability configuration for existing infra nodes deployment\n")
		return runInitConfigExistingNodeHACmd()
	} else {
		msg := "Incorrect argument, please refer help doc."
		return status.Wrap(errors.New(msg), status.ConfigError, UsageTemplate)
	}
}
