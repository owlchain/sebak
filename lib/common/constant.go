package common

import (
	"time"

	"github.com/ulule/limiter"
)

const (
	// BaseFee is the default transaction fee, if fee is lower than BaseFee, the
	// transaction will fail validation.
	BaseFee Amount = 10000

	// BaseReserve is minimum amount of balance for new account. By default, it
	// is `0.1` BOS.
	BaseReserve Amount = 1000000

	// GenesisBlockHeight set the block height of genesis block
	GenesisBlockHeight uint64 = 1

	// GenesisBlockConfirmedTime is the time for the confirmed time of genesis
	// block. This time is of the first commit of SEBAK.
	GenesisBlockConfirmedTime string = "2018-04-17T5:07:31.000000000Z"

	// InflationRatio is the inflation ratio. If the decimal points is over 17,
	// the inflation amount will be 0, considering with `MaximumBalance`. The
	// current value, `0.0000001` will increase `50BOS` in every block(current
	// genesis balance is `5000000000000000`).
	InflationRatio float64 = 0.0000001

	// BlockHeightEndOfInflation sets the block height of inflation end.
	BlockHeightEndOfInflation uint64 = 36000000
	// UnfreezingPeriod is the blocks duration for unfreezing.
	// Frozen account can be unfreezed after passing unfreezing period from unfreezing request.
	// It can be calculated like this. 241920 = 12*60*24*14. This period is considered as about two weeks.
	UnfreezingPeriod uint64 = 241920
)

var (
	// BallotConfirmedTimeAllowDuration is the duration time for ballot from
	// other nodes. If confirmed time of ballot has too late or ahead by
	// BallotConfirmedTimeAllowDuration, it will be considered not-wellformed.
	// For details, `Ballot.IsWellFormed()`
	BallotConfirmedTimeAllowDuration time.Duration = time.Minute * time.Duration(1)

	InflationRatioString string = InflationRatio2String(InflationRatio)

	// RateLimitAPI set the rate limit for API interface, the default value
	// allows 100 requests per minute.
	RateLimitAPI limiter.Rate = limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  100,
	}

	// RateLimitNode set the rate limit for node interface, the default value
	// allows 100 requests per seconds.
	RateLimitNode limiter.Rate = limiter.Rate{
		Period: 1 * time.Second,
		Limit:  100,
	}
)
