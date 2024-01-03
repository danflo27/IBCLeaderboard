package keeper

import (
	"github.com/cosmonaut/leaderboard2/x/leaderboard/types"
)

var _ types.QueryServer = Keeper{}
