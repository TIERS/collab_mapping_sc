// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package controllerswarm

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

const (
	IdxParamAgent            = 0
	IdxParamLength           = 1
	IdxParamMap              = 2
	IdxParamMap1             = 3
	IdxParamMap2             = 4
	IdxParamNoiseThreshold   = 5
	IdxParamWidth            = 6
	IdxParamWindowSize       = 7
	IdxResultConflict        = 8
	IdxResultGlobalMap       = 9
	IdxResultMission         = 10
	IdxStateAgents           = 11
	IdxStateLength           = 12
	IdxStateMissions         = 13
	IdxStateNoiseThreshold   = 14
	IdxStateSubMapAgent      = 15
	IdxStateSubMapByzIndices = 16
	IdxStateSubmittedMaps    = 17
	IdxStateWidth            = 18
	IdxStateWindowSize       = 19
)

const keyMapLen = 20

var keyMap = [keyMapLen]wasmlib.Key{
	ParamAgent,
	ParamLength,
	ParamMap,
	ParamMap1,
	ParamMap2,
	ParamNoiseThreshold,
	ParamWidth,
	ParamWindowSize,
	ResultConflict,
	ResultGlobalMap,
	ResultMission,
	StateAgents,
	StateLength,
	StateMissions,
	StateNoiseThreshold,
	StateSubMapAgent,
	StateSubMapByzIndices,
	StateSubmittedMaps,
	StateWidth,
	StateWindowSize,
}

var idxMap [keyMapLen]wasmlib.Key32
