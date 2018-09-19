package main

import (
	"fmt"
	"sync"

	"github.com/rivine/rivine/modules"
	rivinetypes "github.com/rivine/rivine/types"
	"github.com/threefoldfoundation/tfchain/pkg/persist"
)

// Explorer defines the custom (internal) explorer module,
// used to dump the data of a tfchain network in a meaningful way.
type Explorer struct {
	cs   modules.ConsensusSet
	txdb *persist.TransactionDB

	bcInfo   rivinetypes.BlockchainInfo
	chainCts rivinetypes.ChainConstants

	mut sync.Mutex
}

// NewExplorer creates a new custom intenral explorer module.
// See Explorer for more information.
func NewExplorer(cs modules.ConsensusSet, txdb *persist.TransactionDB, bcInfo rivinetypes.BlockchainInfo, chainCts rivinetypes.ChainConstants, cancel <-chan struct{}) (*Explorer, error) {
	explorer := &Explorer{
		cs:       cs,
		txdb:     txdb,
		bcInfo:   bcInfo,
		chainCts: chainCts,
	}
	err := cs.ConsensusSetSubscribe(explorer, modules.ConsensusChangeBeginning, cancel)
	if err != nil {
		return nil, fmt.Errorf("explorer: failed to subscribe to consensus set: %v", err)
	}
	return explorer, nil
}

// Close the Explorer module.
func (explorer *Explorer) Close() error {
	explorer.mut.Lock()
	defer explorer.mut.Unlock()
	explorer.cs.Unsubscribe(explorer)
	return nil
}

// ProcessConsensusChange implements modules.ConsensusSetSubscriber,
// used to apply/revert blocks to/from our Redis-stored data.
func (explorer *Explorer) ProcessConsensusChange(css modules.ConsensusChange) {
	explorer.mut.Lock()
	defer explorer.mut.Unlock()

	// TODO:
	// * keep track of state
	// * define what to index

	// update reverted blocks
	for _, block := range css.RevertedBlocks {
		// revert miner payouts
		for i, mp := range block.MinerPayouts {
			fmt.Println("revert: minerpayout (type of coinoutput):", i, mp)
		}
		// revert txs
		for i, tx := range block.Transactions {
			fmt.Println("revert: Tx and all its content (depends upon the type):", i, tx)
		}
	}

	// update applied blocks
	for _, block := range css.AppliedBlocks {
		// apply miner payouts
		for i, mp := range block.MinerPayouts {
			fmt.Println("apply: minerpayout (type of coinoutput):", i, mp)
		}
		// apply txs
		for i, tx := range block.Transactions {
			fmt.Println("apply: Tx and all its content (depends upon the type):", i, tx)
		}
	}

	// TODO:
	// update state as well
}
