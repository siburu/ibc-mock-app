package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/spf13/cobra"
)

// NewTxCmd returns the transaction commands for the IBC mockapp module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        "mockapp",
		Short:                      "IBC MockApp transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		NewSendPacketTxCmd(),
	)

	return txCmd
}
