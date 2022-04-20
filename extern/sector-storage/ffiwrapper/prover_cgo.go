//go:build cgo
// +build cgo

package ffiwrapper

import (
	ffi "github.com/filecoin-project/filecoin-ffi"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
)

var ProofProver = proofProver{}

var _ Prover = ProofProver

type proofProver struct{}

func (v proofProver) AggregateSealProofs(aggregateInfo miner.AggregateSealVerifyProofAndInfos, proofs [][]byte) ([]byte, error) {
	return ffi.AggregateSealProofs(aggregateInfo, proofs)
}
