// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package agentswarm
import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type ArrayOfImmutableAgentID struct {
	objID int32
}

func (a ArrayOfImmutableAgentID) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableAgentID) GetAgentID(index int32) wasmlib.ScImmutableAgentID {
	return wasmlib.NewScImmutableAgentID(a.objID, wasmlib.Key32(index))
}

type ArrayOfImmutableOccupancyGrid struct {
	objID int32
}

func (a ArrayOfImmutableOccupancyGrid) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfImmutableOccupancyGrid) GetOccupancyGrid(index int32) ImmutableOccupancyGrid {
	subID := wasmlib.GetObjectID(a.objID, wasmlib.Key32(index), wasmlib.TYPE_ARRAY|wasmlib.TYPE_INT32)
	return ImmutableOccupancyGrid{objID: subID}
}

type ImmutableAgentSwarmState struct {
	id int32
}

func (s ImmutableAgentSwarmState) Controller() wasmlib.ScImmutableChainID {
	return wasmlib.NewScImmutableChainID(s.id, idxMap[IdxStateController])
}

func (s ImmutableAgentSwarmState) MissionID() wasmlib.ScImmutableInt32 {
	return wasmlib.NewScImmutableInt32(s.id, idxMap[IdxStateMissionID])
}

func (s ImmutableAgentSwarmState) Possition() ImmutablePosition {
	return ImmutablePosition{objID: s.id, keyID: idxMap[IdxStatePossition]}
}

func (s ImmutableAgentSwarmState) SubmittedAgent() ArrayOfImmutableAgentID {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateSubmittedAgent], wasmlib.TYPE_ARRAY|wasmlib.TYPE_AGENT_ID)
	return ArrayOfImmutableAgentID{objID: arrID}
}

func (s ImmutableAgentSwarmState) SubmittedMaps() ArrayOfImmutableOccupancyGrid {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateSubmittedMaps], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfImmutableOccupancyGrid{objID: arrID}
}

type ArrayOfMutableAgentID struct {
	objID int32
}

func (a ArrayOfMutableAgentID) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableAgentID) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableAgentID) GetAgentID(index int32) wasmlib.ScMutableAgentID {
	return wasmlib.NewScMutableAgentID(a.objID, wasmlib.Key32(index))
}

type ArrayOfMutableOccupancyGrid struct {
	objID int32
}

func (a ArrayOfMutableOccupancyGrid) Clear() {
	wasmlib.Clear(a.objID)
}

func (a ArrayOfMutableOccupancyGrid) Length() int32 {
	return wasmlib.GetLength(a.objID)
}

func (a ArrayOfMutableOccupancyGrid) GetOccupancyGrid(index int32) MutableOccupancyGrid {
	subID := wasmlib.GetObjectID(a.objID, wasmlib.Key32(index), wasmlib.TYPE_ARRAY|wasmlib.TYPE_INT32)
	return MutableOccupancyGrid{objID: subID}
}

type MutableAgentSwarmState struct {
	id int32
}

func (s MutableAgentSwarmState) Controller() wasmlib.ScMutableChainID {
	return wasmlib.NewScMutableChainID(s.id, idxMap[IdxStateController])
}

func (s MutableAgentSwarmState) MissionID() wasmlib.ScMutableInt32 {
	return wasmlib.NewScMutableInt32(s.id, idxMap[IdxStateMissionID])
}

func (s MutableAgentSwarmState) Possition() MutablePosition {
	return MutablePosition{objID: s.id, keyID: idxMap[IdxStatePossition]}
}

func (s MutableAgentSwarmState) SubmittedAgent() ArrayOfMutableAgentID {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateSubmittedAgent], wasmlib.TYPE_ARRAY|wasmlib.TYPE_AGENT_ID)
	return ArrayOfMutableAgentID{objID: arrID}
}

func (s MutableAgentSwarmState) SubmittedMaps() ArrayOfMutableOccupancyGrid {
	arrID := wasmlib.GetObjectID(s.id, idxMap[IdxStateSubmittedMaps], wasmlib.TYPE_ARRAY|wasmlib.TYPE_BYTES)
	return ArrayOfMutableOccupancyGrid{objID: arrID}
}
