package resource

import (
	"strings"

	"boscoin.io/sebak/lib/block"
	"boscoin.io/sebak/lib/common"
	"github.com/nvellon/hal"
)

type FrozenAccountState string

const (
	FrozenState    FrozenAccountState = "frozen"
	MeltingState   FrozenAccountState = "melting"
	UnfreezedState FrozenAccountState = "unfreezed"
	RefundState    FrozenAccountState = "refund"
)

type FrozenAccount struct {
	ba                           *block.BlockAccount
	createdBlockheight           uint64
	createdOpHash                string
	createdSequenceId            uint64
	initialAmount                common.Amount
	freezingState                FrozenAccountState
	unfreezingRequestBlockheight uint64
	unfreezingRequestOpHash      string
	paymentOpHash                string
}

func NewFrozenAccount(ba *block.BlockAccount, create_block_height uint64, create_opHash string, sequenceid uint64, amount common.Amount, state FrozenAccountState, unfreezing_block_height uint64, unfreezing_opHash string, payment_opHash string) *FrozenAccount {
	fa := &FrozenAccount{
		ba:                           ba,
		createdBlockheight:           create_block_height,
		createdOpHash:                create_opHash,
		createdSequenceId:            sequenceid,
		initialAmount:                amount,
		freezingState:                state,
		unfreezingRequestBlockheight: unfreezing_block_height,
		unfreezingRequestOpHash:      unfreezing_opHash,
		paymentOpHash:                payment_opHash,
	}
	return fa
}

func (fa FrozenAccount) GetMap() hal.Entry {
	return hal.Entry{
		"address":                 fa.ba.Address,
		"linked":                  fa.ba.Linked,
		"create_block_height":     fa.createdBlockheight,
		"create_opHash":           fa.createdOpHash,
		"sequence_id":             fa.createdSequenceId,
		"amount":                  fa.initialAmount,
		"state":                   fa.freezingState,
		"unfreezing_block_height": fa.unfreezingRequestBlockheight,
		"unfreezing_OpHash":       fa.unfreezingRequestOpHash,
		"payment_opHash":          fa.paymentOpHash,
	}
}

func (fa FrozenAccount) Resource() *hal.Resource {
	r := hal.NewResource(fa, fa.LinkSelf())

	return r
}

func (fa FrozenAccount) LinkSelf() string {
	address := fa.ba.Linked

	return strings.Replace(URLAccountFrozenAccounts, "{id}", address, -1)
}
