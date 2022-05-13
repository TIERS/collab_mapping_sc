// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package controllerswarm

import (
	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
)

func funcRegister(ctx wasmlib.ScFuncContext, f *RegisterContext) {
	agents := f.State.Agents()
	newAgent := f.Params.Agent().Value()

	// Check if agent is already registerd or not
	value := agents.GetInt32(newAgent)
	ctx.Require(!value.Exists(), "Agent is already registerd.")

	// 0: Benign Agent
	value.SetValue(0)
}

func funcRequestMission(ctx wasmlib.ScFuncContext, f *RequestMissionContext) {
	agent := f.Params.Agent().Value()
	agentStatus := f.State.Agents().GetInt32(agent)

	// Check if agent swarm is registerd or not
	ctx.Require(agentStatus.Exists(), "Agent is not registered.")
	// Check if agent swarm is not detected as a byzantine agent
	ctx.Require(agentStatus.Value() == 0, "Agent is marked as byzantine!")

	// return mission id
	missionID := f.State.Missions().Value() + 1
	f.State.Missions().SetValue(missionID)
}

func funcSubmitMap(ctx wasmlib.ScFuncContext, f *SubmitMapContext) {
	// Check if the length and width of the submitted map is correct
	lenMap := f.Params.Map().Length()
	ctx.Require(lenMap == f.State.Width().Value()*f.State.Length().Value(), "Wrong map!")

	agent := f.Params.Agent().Value() // agent submitting the map
	mapInd := f.State.SubmittedMaps().Length()
	var mapByzInd MapByzIndex // map Byzantine Index
	mapByzInd.Conflict = 0
	mapByzInd.NonConflict = 0

	// Iterate over all the submitted maps and compare it with the new one
	for i := int32(0); i < mapInd; i++ {
		compareFunc := ScFuncs.CompareMap(ctx)

		// copy input map to the map1 input of the compare function
		for j := int32(0); j < lenMap; j++ {
			v := f.Params.Map().GetInt32(j).Value()
			compareFunc.Params.Map2().GetInt32(j).SetValue(v)
		}

		// copy current map to the map2 input of the compare function
		for j := int32(0); j < lenMap; j++ {
			v := f.State.SubmittedMaps().GetOccupancyGrid(i).GetInt32(j).Value()
			compareFunc.Params.Map1().GetInt32(j).SetValue(v)
		}

		// call function and read the result
		compareFunc.Func.Call()
		conflict := compareFunc.Results.Conflict().Value()

		// A ref to the current maps byzIndex
		curByzInd := f.State.SubMapByzIndices().GetMapByzIndex(i).Value()

		// Update values
		if conflict == int32(0) {
			mapByzInd.NonConflict += 1
			curByzInd.NonConflict += 1
		} else {
			mapByzInd.Conflict += 1
			curByzInd.Conflict += 1
		}

		// Save the updated current maps byzIndex
		f.State.SubMapByzIndices().GetMapByzIndex(i).SetValue(curByzInd)

		// Update the current map's agent status
		agent_ := f.State.SubMapAgent().GetAgentID(i).Value()
		if curByzInd.Conflict > curByzInd.NonConflict {
			f.State.Agents().GetInt32(agent_).SetValue(1) // Byzantine
		}
	}

	// add to submitted maps
	f.State.SubMapByzIndices().GetMapByzIndex(mapInd).SetValue(&mapByzInd)
	for i := int32(0); i < lenMap; i++ {
		v := f.Params.Map().GetInt32(i).Value()
		f.State.SubmittedMaps().GetOccupancyGrid(mapInd).GetInt32(i).SetValue(v)
	}
	f.State.SubMapAgent().GetAgentID(mapInd).SetValue(agent)

	// Submitting agent's status
	if mapByzInd.Conflict > mapByzInd.NonConflict {
		f.State.Agents().GetInt32(agent).SetValue(1) // Byzantine
	}
}

func viewGlobalMap(ctx wasmlib.ScViewContext, f *GlobalMapContext) {
	len := f.State.Width().Value() * f.State.Length().Value()
	// Initializing resulting map with 0 values
	var res []int32
	for i := int32(0); i < len; i++ {
		res[i] = 0
	}

	// Iterate over all submitted maps to merge them
	for i := int32(0); i < f.State.SubmittedMaps().Length(); i++ {
		// Check if the current map's agent is byzantine or not
		agent := f.State.SubMapAgent().GetAgentID(i).Value()
		if f.State.Agents().GetInt32(agent).Value() == 1 {
			continue
		}

		// Iterate over each cell of map.
		// Add the value of each cell to result map
		for j := int32(0); j < len; j++ {
			res[j] = res[j] + f.State.SubmittedMaps().GetOccupancyGrid(i).GetInt32(j).Value()
		}
	}

	// return the merged map
	for i := int32(0); i < len; i++ {
		if res[i] == 0 {
			f.Results.GlobalMap().GetInt32(i).SetValue(0)
		} else if res[i] > 0 {
			f.Results.GlobalMap().GetInt32(i).SetValue(1)
		} else if res[i] < 0 {
			f.Results.GlobalMap().GetInt32(i).SetValue(-1)
		}
	}
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	length := f.Params.Length().Value()
	width := f.Params.Width().Value()
	windowSize := f.Params.WindowSize().Value()
	noiseThreshold := f.Params.NoiseThreshold().Value()

	f.State.Length().SetValue(length)
	f.State.Width().SetValue(width)
	f.State.WindowSize().SetValue(windowSize)
	f.State.Missions().SetValue(0)
	f.State.NoiseThreshold().SetValue(noiseThreshold)
}

func viewGetMission(ctx wasmlib.ScViewContext, f *GetMissionContext) {
	m := f.State.Missions().Value()
	f.Results.Mission().SetValue(m)
}

func funcCompareMap(ctx wasmlib.ScFuncContext, f *CompareMapContext) {
	m1 := f.Params.Map1()
	m2 := f.Params.Map2()
	ws := f.State.WindowSize().Value()
	L := f.State.Length().Value()
	W := f.State.Width().Value()
	nT := f.State.NoiseThreshold().Value() * ws * ws / 100

	byz := false
	// Iterating over each windows (l, w)
	for l := int32(0); l < L/ws; l++ {
		for w := int32(0); w < W/ws; w++ {

			similarity := int32(0)
			diff := int32(0)
			z1 := int32(0)
			z2 := int32(0)

			// Iterating over each row of window
			for l_ := int32(0); l_ < ws; l_++ {
				start := (l*ws+l_)*W + w*ws
				end := (l*ws+l_)*W + (w+1)*ws

				for k := start; k < end; k++ {
					// values of each cell in row
					v1 := m1.GetInt32(k).Value()
					v2 := m2.GetInt32(k).Value()

					// simalirities
					if v1 == v2 && v1 != int32(0) {
						similarity += 1
					}

					// diffs
					if v1 != v2 && v1 != int32(0) && v2 != int32(0) {
						diff += 1
					}

					// zeros in map1
					if v1 == int32(0) {
						z1 += 1
					}

					// zeros in map2
					if v2 == int32(0) {
						z2 += 1
					}
				}
			} // End of the iteration over rows

			// Detect az conflicting
			if (diff+similarity != int32(0)) || (z1 > nT) || (z2 > nT) {
				byz = true
			}
		} // End of the iteration over windows
	}

	// Return 1 if detected something
	if byz {
		f.Results.Conflict().SetValue(1)
	} else {
		f.Results.Conflict().SetValue(0)
	}
}
