package root

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/gauravgahlot/syncroot/cmd/server"
	"github.com/gauravgahlot/syncroot/internal/config"
)

func Command(log *zap.Logger, cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "syncer",
		Short:        "syncer CLI",
		SilenceUsage: true,
	}

	cmd.AddCommand(server.Command(log, cfg))

	return cmd
}
