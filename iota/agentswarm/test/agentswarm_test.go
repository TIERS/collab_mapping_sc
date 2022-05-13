package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/wasm/agentswarm/go/agentswarm"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, agentswarm.ScName, agentswarm.OnLoad)
	require.NoError(t, ctx.ContractExists(agentswarm.ScName))
}
