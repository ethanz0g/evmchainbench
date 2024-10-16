package cmd

import (
	"fmt"

	"github.com/0glabs/evmchainbench/cmd/option"
	"github.com/0glabs/evmchainbench/lib/cmd/gentx"
	"github.com/spf13/cobra"
)

var gentxCmd = &cobra.Command{
	Use:   "gentx",
	Short: "To generate transactions and store them onto disk",
	Long:  "To generate transactions and store them onto disk",
	Run: func(cmd *cobra.Command, args []string) {

		httpRpc, _ := cmd.Flags().GetString("http-rpc")
		faucetPrivateKey, _ := cmd.Flags().GetString("faucet-private-key")
		senderCount, _ := cmd.Flags().GetInt("sender-count")
		txCount, _ := cmd.Flags().GetInt("tx-count")
		txStoreDir, _ := cmd.Flags().GetString("tx-store-dir")
		mempool, _ := cmd.Flags().GetInt("mempool")

		gentx.GenTx(httpRpc, faucetPrivateKey, senderCount, txCount, txStoreDir, mempool)
		fmt.Println("gentx called")
	},
}

func init() {
	rootCmd.AddCommand(gentxCmd)
	option.OptionsForGeneration(gentxCmd)
	option.OptionsForTxStore(gentxCmd)
}
