/*
Copyright © 2026 JACOB ARTHURS
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	flagName      string
	flagArtifacts []string
	flagVersion   string
	flagSummary   string
	flagLicense   string
	flagDryRun    bool
	flagReadme    string
	flagNoReadme  bool
)

var rootCmd = &cobra.Command{
	Use:          "shipbin",
	SilenceUsage: true,
	Short:        "Ship binaries to npm and PyPI",
	Long: `Publishes pre-built binaries to npm and PyPI.

Assembles platform-specific packages from the provided artifacts,
then publishes them to the target registry.`,
	Example: `  # Publish to npm
  shipbin npm --org myorg --name mytool \
    --artifact linux/amd64:dist/mytool_linux_amd64/mytool \
    --artifact darwin/arm64:dist/mytool_darwin_arm64/mytool \
    --artifact windows/amd64:dist/mytool_windows_amd64/mytool.exe

  # Publish to PyPI
  shipbin pypi --name mytool \
    --artifact linux/amd64:dist/mytool_linux_amd64/mytool \
    --artifact darwin/arm64:dist/mytool_darwin_arm64/mytool \
    --artifact windows/amd64:dist/mytool_windows_amd64/mytool.exe`,
}

func init() {
	rootCmd.PersistentFlags().StringVar(&flagName, "name", "", "binary name")
	rootCmd.PersistentFlags().StringArrayVar(&flagArtifacts, "artifact", nil, "os/arch:path mapping (repeatable)")
	rootCmd.PersistentFlags().StringVar(&flagVersion, "version", "", "release version (defaults to current git tag)")
	rootCmd.PersistentFlags().StringVar(&flagSummary, "summary", "", "short description of the package (optional)")
	rootCmd.PersistentFlags().StringVar(&flagLicense, "license", "", "license identifier (e.g. MIT, Apache-2.0)")
	rootCmd.PersistentFlags().BoolVar(&flagDryRun, "dry-run", false, "print what would be published without publishing")
	rootCmd.PersistentFlags().StringVar(&flagReadme, "readme", "", "path to README file (auto-detects README.md/.rst/.txt in current directory if not set)")
	rootCmd.PersistentFlags().BoolVar(&flagNoReadme, "no-readme", false, "disable README auto-detection and omit readme from the package")

	if err := rootCmd.MarkPersistentFlagRequired("name"); err != nil {
		panic(err)
	}
	if err := rootCmd.MarkPersistentFlagRequired("artifact"); err != nil {
		panic(err)
	}

	rootCmd.AddCommand(npmCmd)
	rootCmd.AddCommand(pypiCmd)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
