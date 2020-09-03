package sealing

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin/miner"
)

// Epochs
const SealRandomnessLookback = miner.ChainFinality

// Epochs
func SealRandomnessLookbackLimit() abi.ChainEpoch {
	return miner.MaxPreCommitRandomnessLookback
}

// Epochs
const InteractivePoRepConfidence = 6
