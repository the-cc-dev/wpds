package commands

import (
	"log"
	"os"
	"runtime/pprof"

	"github.com/peterbooker/wpds/internal/pkg/config"
	"github.com/peterbooker/wpds/internal/pkg/context"
	"github.com/peterbooker/wpds/internal/pkg/slurper"
	"github.com/peterbooker/wpds/internal/pkg/stats"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.AddCommand(pluginsUpdateCmd)
	updateCmd.AddCommand(themesUpdateCmd)
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update files from the WordPress Directory.",
	Long:  `Update Plugin or Theme files from their WordPress Directory.`,
}

var pluginsUpdateCmd = &cobra.Command{
	Use:     "plugins",
	Short:   "Update Plugin files.",
	Long:    ``,
	Example: `wpds update plugins -c 250`,
	Run: func(cmd *cobra.Command, args []string) {

		if CPUProf != "" {
			f, err := os.Create(CPUProf)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		if (C < 10) || (C > 10000) {
			log.Printf("Flag (concurrent-actions, c) out of permitted range (10-10000).\n")
			os.Exit(1)
		}

		// Get Config Details
		name := config.GetName()
		version := config.GetVersion()

		// Check if SVN is installed
		// Used if available, as it is more reliable than the HTTP API
		connector := "api"
		if slurper.CheckForSVN() {
			connector = "svn"
		}

		// Get Working Directory
		wd, _ := os.Getwd()

		// Create new Stats
		stats := stats.New()

		ctx := &context.Context{
			Name:              name,
			Version:           version,
			ConcurrentActions: C,
			ExtensionType:     "plugins",
			FileType:          F,
			Connector:         connector,
			CurrentRevision:   0,
			LatestRevision:    0,
			WorkingDirectory:  wd,
			Stats:             stats,
		}

		log.Println("Updating Plugins...")

		slurper.StartUpdate(ctx)

	},
}

var themesUpdateCmd = &cobra.Command{
	Use:     "themes",
	Short:   "Update Theme files.",
	Long:    ``,
	Example: `wpds update themes -c 250`,
	Run: func(cmd *cobra.Command, args []string) {

		if CPUProf != "" {
			f, err := os.Create(CPUProf)
			if err != nil {
				log.Fatal(err)
			}
			pprof.StartCPUProfile(f)
			defer pprof.StopCPUProfile()
		}

		if (C < 10) || (C > 10000) {
			log.Printf("Flag (concurrent-actions, c) out of permitted range (10-1000).\n")
			os.Exit(1)
		}

		// Get Config Details
		name := config.GetName()
		version := config.GetVersion()

		// Check if SVN is installed
		// Used if available, as it is more reliable than the HTTP API
		connector := "api"
		if slurper.CheckForSVN() {
			connector = "svn"
		}

		// Get Working Directory
		wd, _ := os.Getwd()

		// Create new Stats
		stats := stats.New()

		ctx := &context.Context{
			Name:              name,
			Version:           version,
			ConcurrentActions: C,
			ExtensionType:     "themes",
			FileType:          F,
			Connector:         connector,
			CurrentRevision:   0,
			LatestRevision:    0,
			WorkingDirectory:  wd,
			Stats:             stats,
		}

		log.Println("Updating Themes...")

		slurper.StartUpdate(ctx)

	},
}
