// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package agentswarm

import (
	"github.com/iotaledger/wasp/contracts/wasm/controllerswarm/go/controllerswarm"
	"github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"
)

func funcSubmitLastMap(ctx wasmlib.ScFuncContext, f *SubmitLastMapContext) {
	len := f.State.SubmittedMaps().Length()
	agent := f.State.SubmittedAgent().GetAgentID(len - 1).Value()

	submitFunc := controllerswarm.ScFuncs.SubmitMap(ctx)
	submitFunc.Params.Agent().SetValue(agent)
	for i := int32(0); i < f.State.SubmittedMaps().GetOccupancyGrid(len-1).Length(); i++ {
		val := f.State.SubmittedMaps().GetOccupancyGrid(len - 1).GetInt32(i).Value()
		submitFunc.Params.Map().GetInt32(i).SetValue(val)
	}
}

func funcUploadMap(ctx wasmlib.ScFuncContext, f *UploadMapContext) {
	lenAgents := f.State.SubmittedAgent().Length()
	agent := f.Params.Agent().Value()
	f.State.SubmittedAgent().GetAgentID(lenAgents).SetValue(agent)

	lenMaps := f.State.SubmittedMaps().Length()
	for i := int32(0); i < f.Params.Map().Length(); i++ {
		val := f.Params.Map().GetInt32(i).Value()
		f.State.SubmittedMaps().GetOccupancyGrid(lenMaps).GetInt32(i).SetValue(val)
	}
}

func funcRegister(ctx wasmlib.ScFuncContext, f *RegisterContext) {
	registerFunc := controllerswarm.ScFuncs.Register(ctx)
	registerFunc.Params.Agent().SetValue(f.Params.Agent().Value())
	registerFunc.Func.TransferIotas(1).PostToChain(f.State.Controller().Value())
}

func funcInit(ctx wasmlib.ScFuncContext, f *InitContext) {
	f.State.Controller().SetValue(f.Params.Controller().Value())
}

func funcRequestMission(ctx wasmlib.ScFuncContext, f *RequestMissionContext) {
	reqFunc := controllerswarm.ScFuncs.RequestMission(ctx)
	reqFunc.Params.Agent().SetValue(f.Params.Agent().Value())
	reqFunc.Func.TransferIotas(1).PostToChain(f.State.Controller().Value())
}
