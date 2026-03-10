/*
Copyright © 2026 JACOB ARTHURS
*/
package cmd

import (
	"github.com/jacobarthurs/shipbin/internal/config"
	"github.com/jacobarthurs/shipbin/internal/pypi"
	"github.com/spf13/cobra"
)

var pypiCmd = &cobra.Command{
	Use:   "pypi",
	Short: "Publish binaries to PyPI",
	Long: `Publishes pre-built binaries to PyPI.

Builds a platform-specific wheel for each artifact containing the binary and a
Python shim that locates and executes it. Users install the package with pip and
the correct wheel is resolved automatically based on their platform.`,
	Example: `  # Publish pre-built binaries
  shipbin pypi --name mytool --version v1.2.3 \
    --artifact linux/amd64:dist/mytool_linux_amd64/mytool \
    --artifact darwin/amd64:dist/mytool_darwin_amd64/mytool \
    --artifact windows/amd64:dist/mytool_windows_amd64/mytool.exe

  # Dry run to preview without publishing
  shipbin pypi --name mytool --dry-run \
    --artifact linux/amd64:dist/mytool_linux_amd64/mytool

  # Publish with a license and summary
  shipbin pypi --name mytool --license MIT --summary "A useful CLI tool" \
    --artifact linux/amd64:dist/mytool_linux_amd64/mytool`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg, err := buildPypiConfig()
		if err != nil {
			return err
		}
		return pypi.Publish(cfg)
	},
}

func buildPypiConfig() (*pypi.Config, error) {
	version, err := config.ResolveVersion(flagVersion)
	if err != nil {
		return nil, err
	}

	artifacts, err := config.ParseArtifacts(flagArtifacts)
	if err != nil {
		return nil, err
	}

	cfg := &pypi.Config{
		Name:      flagName,
		Version:   version,
		Artifacts: artifacts,
		Summary:   flagSummary,
		License:   flagLicense,
		Readme:    config.ResolveReadme(flagReadme, flagNoReadme),
		DryRun:    flagDryRun,
	}

	return cfg, nil
}

func init() {}
