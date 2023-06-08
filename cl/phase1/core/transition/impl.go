package transition

import (
	"fmt"

	"github.com/Giulio2002/bls"
	"github.com/ledgerwatch/erigon/cl/cltypes"
	"github.com/ledgerwatch/erigon/cl/fork"
	"github.com/ledgerwatch/erigon/cl/phase1/core/state"
	state2 "github.com/ledgerwatch/erigon/cl/phase1/core/state"
	"github.com/ledgerwatch/erigon/cl/phase1/core/transition/machine"
)

var _ machine.Interface = (*impl)(nil)

var DefaultMachine = &impl{}
var ValidatingMachine = &impl{FullValidation: true}

type impl struct {
	FullValidation bool
}

func (I *impl) VerifyTransition(s *state.BeaconState, currentBlock *cltypes.BeaconBlock) error {
	if !I.FullValidation {
		return nil
	}
	expectedStateRoot, err := s.HashSSZ()
	if err != nil {
		return fmt.Errorf("unable to generate state root: %v", err)
	}
	if expectedStateRoot != currentBlock.StateRoot {
		return fmt.Errorf("expected state root differs from received state root")
	}
	return nil
}

func (I *impl) VerifyBlockSignature(s *state.BeaconState, block *cltypes.SignedBeaconBlock) error {
	if !I.FullValidation {
		return nil
	}
	valid, err := verifyBlockSignature(s, block)
	if err != nil {
		return fmt.Errorf("error validating block signature: %v", err)
	}
	if !valid {
		return fmt.Errorf("block not valid")
	}
	return nil
}

func verifyBlockSignature(s *state2.BeaconState, block *cltypes.SignedBeaconBlock) (bool, error) {
	proposer, err := s.ValidatorForValidatorIndex(int(block.Block.ProposerIndex))
	if err != nil {
		return false, err
	}
	domain, err := s.GetDomain(s.BeaconConfig().DomainBeaconProposer, state2.Epoch(s.BeaconState))
	if err != nil {
		return false, err
	}
	sigRoot, err := fork.ComputeSigningRoot(block.Block, domain)
	if err != nil {
		return false, err
	}
	pk := proposer.PublicKey()
	return bls.Verify(block.Signature[:], sigRoot[:], pk[:])
}
