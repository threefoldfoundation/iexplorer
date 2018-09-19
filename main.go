package main

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/threefoldfoundation/tfchain/pkg/config"
)

func main() {
	cmd := new(Commands)
	cmd.BlockchainInfo = config.GetBlockchainInfo()

	// define commands
	cmdRoot := &cobra.Command{
		Use:   "iexplorer",
		Short: "start the iexplorer daemon",
		Args:  cobra.ExactArgs(0),
		RunE:  cmd.Root,
	}

	cmdVersion := &cobra.Command{
		Use:   "version",
		Short: "show versions of this tool",
		Args:  cobra.ExactArgs(0),
		Run:   cmd.Version,
	}

	// define command tree
	cmdRoot.AddCommand(
		cmdVersion,
	)

	// define flags
	cmdRoot.Flags().StringVarP(
		&cmd.RootPersistentDir,
		"persistent-directory", "d",
		cmd.RootPersistentDir,
		"location of the root diretory used to store persistent data of the daemon of "+cmd.BlockchainInfo.Name,
	)
	/*
		TODO: define any flags required to be able to connect
		to the database of choice
	*/
	cmdRoot.Flags().StringVar(
		&cmd.ProfilingAddr,
		"profile-addr",
		cmd.ProfilingAddr,
		"enables profiling of this iexplorer instance as an http service",
	)
	cmdRoot.Flags().StringVarP(
		&cmd.BlockchainInfo.NetworkName,
		"network", "n",
		cmd.BlockchainInfo.NetworkName,
		"the name of the network to which the daemon connects, one of {standard,testnet}",
	)

	// execute logic
	if err := cmdRoot.Execute(); err != nil {
		os.Exit(1)
	}
}
