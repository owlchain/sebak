package errors

import "errors"

var New = errors.New

var (
	ErrorBlockAlreadyExists                        = NewError(100, "already exists in block")
	ErrorHashDoesNotMatch                          = NewError(101, "`Hash` does not match")
	ErrorSignatureVerificationFailed               = NewError(102, "signature verification failed")
	ErrorBadPublicAddress                          = NewError(103, "failed to parse public address")
	ErrorInvalidFee                                = NewError(104, "invalid fee")
	ErrorInvalidOperation                          = NewError(105, "invalid operation")
	ErrorNewButKnownMessage                        = NewError(106, "received new, but known message")
	ErrorInvalidState                              = NewError(107, "found invalid state")
	ErrorInvalidVotingThresholdPolicy              = NewError(108, "invalid `VotingThresholdPolicy`")
	ErrorBallotEmptyMessage                        = NewError(109, "init state ballot does not have `Message`")
	ErrorInvalidHash                               = NewError(110, "invalid `Hash`")
	ErrorInvalidMessage                            = NewError(111, "invalid `Message`")
	ErrorBallotHasMessage                          = NewError(112, "none-init state ballot must not have `Message`")
	ErrorVotingResultAlreadyExists                 = NewError(113, "`VotingResult` already exists")
	ErrorVotingResultNotFound                      = NewError(114, "`VotingResult` not found")
	ErrorVotingResultFailedToSetState              = NewError(115, "failed to set the new state to `VotingResult`")
	ErrorVotingResultNotInBox                      = NewError(116, "ballot is not in here")
	ErrorBallotNoVoting                            = NewError(118, "ballot has no `Voting`")
	ErrorBallotNoNodeKey                           = NewError(119, "ballot has no `NodeKey`")
	ErrorVotingThresholdInvalidValidators          = NewError(120, "invalid validators")
	ErrorBallotHasInvalidState                     = NewError(121, "ballot has invalid state")
	ErrorVotingResultFailedToClose                 = NewError(122, "failed to close `VotingResult`")
	ErrorTransactionEmptyOperations                = NewError(123, "operations needs in transaction")
	ErrorAlreadySaved                              = NewError(124, "already saved")
	ErrorDuplicatedOperation                       = NewError(125, "duplicated operations in transaction")
	ErrorUnknownOperationType                      = NewError(126, "unknown operation type")
	ErrorTypeOperationBodyNotMatched               = NewError(127, "operation type and it's type does not match")
	ErrorBlockAccountDoesNotExists                 = NewError(128, "account does not exists in block")
	ErrorBlockAccountAlreadyExists                 = NewError(129, "account already exists in block")
	ErrorAccountBalanceUnderZero                   = NewError(130, "account balance will be under zero")
	ErrorMaximumBalanceReached                     = NewError(131, "monetary amount would be greater than the total supply of coins")
	ErrorStorageRecordDoesNotExist                 = NewError(132, "record does not exist in storage")
	ErrorTransactionInvalidSequenceID              = NewError(133, "invalid sequenceID found")
	ErrorBlockTransactionDoesNotExists             = NewError(134, "transaction does not exists in block")
	ErrorBlockOperationDoesNotExists               = NewError(135, "operation does not exists in block")
	ErrorRoundVoteNotFound                         = NewError(136, "`RoundVote` not found")
	ErrorBlockNotFound                             = NewError(137, "Block not found")
	ErrorTransactionExcessAbilityToPay             = NewError(138, "Transaction requests over ability to pay")
	ErrorTransactionSameSource                     = NewError(139, "Same transaction source found in ballot")
	ErrorTransactionNotFound                       = NewError(140, "Transaction not found")
	ErrorBallotFromUnknownValidator                = NewError(141, "ballot from unknown validator")
	ErrorBallotAlreadyFinished                     = NewError(142, "ballot already finished")
	ErrorBallotAlreadyVoted                        = NewError(143, "ballot already voted")
	ErrorBallotHasOverMaxTransactionsInBallot      = NewError(144, "too many transactions in ballot")
	ErrorMessageHasIncorrectTime                   = NewError(145, "time in message is not correct")
	ErrorInvalidQueryString                        = NewError(146, "found invalid query string")
	ErrorInvalidContentType                        = NewError(147, "found invalid 'Content-Type'")
	ErrorStorageRecordAlreadyExists                = NewError(148, "record already exists in storage")
	ErrorStorageCoreError                          = NewError(149, "storage error")
	ErrorContentTypeNotJSON                        = NewError(150, "`Content-Type` must be 'application/json'")
	ErrorTransactionHasOverMaxOperations           = NewError(151, "too many operations in transaction")
	ErrorOperationAmountUnderflow                  = NewError(152, "invalid `Amount`: lower than 1")
	ErrorFrozenAccountNoDeposit                    = NewError(153, "frozen account can not receive payment")
	ErrorFrozenAccountCreationWholeUnit            = NewError(154, "frozen account balance must be a whole number of units (10k)")
	ErrorFrozenAccountMustWithdrawEverything       = NewError(155, "frozen account can only withdraw the full amount (minus tx fee)")
	ErrorInsufficientAmountNewAccount              = NewError(156, "insufficient amount for new account")
	ErrorOperationBodyInsufficient                 = NewError(157, "operation body insufficient")
	ErrorOperationAmountOverflow                   = NewError(158, "invalid `Amount`: over than expected")
	ErrorWrongBlockFound                           = NewError(159, "wrong Block found")
	ErrorInvalidProposerTransaction                = NewError(160, "invalid proposer transaction found")
	ErrorInvalidInflationRatio                     = NewError(161, "invalid inflation ratio found")
	ErrorNotImplemented                            = NewError(162, "not implemented")
	ErrorHTTPProblem                               = NewError(163, "http failed to get response")
	ErrorInvalidTransaction                        = NewError(164, "invalid transaction")
	ErrorNotMatcHTTPRouter                         = NewError(165, "doesn't match http router")
	ErrorUnfreezingFromInvalidAccount              = NewError(166, "unfreezing should be done from a frozen account")
	ErrorUnfreezingToInvalidLinkedAccount          = NewError(167, "unfreezing should be done to a valid linked account")
	ErrorUnfreezingNotReachedExpiration            = NewError(168, "unfreezing should pass 241920 blockheight from unfreezing request")
	ErrorFrozenAccountMustCreatedFromLinkedAccount = NewError(169, "frozen account create-transaction must be generated from the linked account")
	ErrorUnfreezingRequestNotRequested             = NewError(170, "unfreezing must be generated after the unfreezing request")
	ErrorUnfreezingRequestAlreadyReceived          = NewError(171, "unfreezing request already received from client")
	ErrorTooManyRequests                           = NewError(172, "too many requests; reached limit")
	ErrorHTTPServerError                           = NewError(173, "Internal Server Error")
)
