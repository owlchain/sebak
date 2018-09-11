package common

const (
	BlockPrefixHash                       = string(0x00)
	BlockPrefixConfirmed                  = string(0x01)

	BlockTransactionPrefixHash            = string(0x10)
	BlockTransactionPrefixSource          = string(0x11)
	BlockTransactionPrefixConfirmed       = string(0x12)
	BlockTransactionPrefixAccount         = string(0x13)
	BlockTransactionPrefixBlock           = string(0x14)

	BlockOperationPrefixHash              = string(0x20)
	BlockOperationPrefixTxHash            = string(0x21)
	BlockOperationPrefixSource            = string(0x22)
	BlockOperationPrefixTarget            = string(0x23)
	BlockOperationPrefixPeers             = string(0x24)

	BlockAccountPrefixAddress             = string(0x30)
	BlockAccountPrefixCreated             = string(0x31)
	BlockAccountSequenceIDPrefix          = string(0x32)
	BlockAccountSequenceIDByAddressPrefix = string(0x33)
)
