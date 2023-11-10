package etherman

import (
	"context"
	"log"
	"math/big"
	"sort"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

type ConsensusMember struct {
	Addr common.Address
	URL  string
}

type ConsensusMembers []ConsensusMember

func (s ConsensusMembers) Len() int { return len(s) }
func (s ConsensusMembers) Less(i, j int) bool {
	return strings.ToUpper(s[i].Addr.Hex()) < strings.ToUpper(s[j].Addr.Hex())
}
func (s ConsensusMembers) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

type Consensus struct {
	requiredSignatures *big.Int
	members            ConsensusMembers
}

func NewConsensus(requiredSignatures *big.Int, members ConsensusMembers) *Consensus {
	sort.Sort(members)
	return &Consensus{
		requiredSignatures: requiredSignatures,
		members:            members,
	}
}

// SetupConsensus call the crater contract to setup the consensus members.
// Synchronous, the func will only return after the tx is mined
func (e *Etherman) SetupConsensus(consensus Consensus) (err error) {
	ctx := context.TODO()
	addrsBytes := []byte{}
	urls := []string{}
	for _, member := range consensus.members {
		addrsBytes = append(addrsBytes, member.Addr.Bytes()...)
		urls = append(urls, member.URL)
	}

	tx, err := e.contract.SetupCommittee(
		e.auth,
		consensus.requiredSignatures,
		urls,
		addrsBytes,
	)
	if err != nil {
		return err
	}
	log.Println("Sending tx with a timeout of 3 mins")
	return e.WaitTxToBeMined(ctx, tx, time.Minute*3)
}
