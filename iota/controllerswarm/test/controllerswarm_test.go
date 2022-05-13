package test

import (
	"testing"

	"github.com/iotaledger/wasp/contracts/wasm/controllerswarm/go/controllerswarm"
	"github.com/iotaledger/wasp/packages/vm/wasmsolo"
	"github.com/stretchr/testify/require"
)

func TestDeploy(t *testing.T) {
	ctx := wasmsolo.NewSoloContext(t, controllerswarm.ScName, controllerswarm.OnLoad)
	require.NoError(t, ctx.ContractExists(controllerswarm.ScName))
}
