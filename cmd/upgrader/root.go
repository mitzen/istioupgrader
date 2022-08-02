package upgrader

import (
	"fmt"
	"os"

	"github.com/mitzen/istioupgrader/pkg/feature"
	"github.com/spf13/cobra"
)

// Testing
// Example of use case
// istioupgrader.exe --type=canary --version=10.20.30

func Execute() {

	var istioUpgradeType string
	var targetVersion string

	var versionCmd = &cobra.Command{
		Use:   "version",
		Short: "Get version info",
		Long:  `Get version info`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("version 1.0")
		},
	}

	var rootCmd = &cobra.Command{
		Use:   "upgrade",
		Short: "upgrade",
		Long:  `Upgrade like the pro and do inplace or canary without breaking a sweat!`,
		Run: func(cmd *cobra.Command, args []string) {

			fmt.Printf("%s\n", istioUpgradeType)
			fmt.Printf("%s\n", targetVersion)

			istioUpgrader := feature.IstioUpgrade{UpgradeType: istioUpgradeType,
				VersionSelected: targetVersion, Cmd: cmd}
			istioUpgrader.Execute()
		},
	}

	rootCmd.Flags().StringVarP(&istioUpgradeType, "type", "t", "inplace", "Upgrade type")
	rootCmd.MarkFlagRequired("type")
	rootCmd.Flags().StringVarP(&targetVersion, "version", "v", "1.0.0", "Target version to upgrade")
	rootCmd.MarkFlagRequired("version")

	rootCmd.AddCommand(versionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
