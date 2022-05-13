// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package agentswarm

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type InitCall struct {
	Func   *wasmlib.ScInitFunc
	Params MutableInitParams
}

type RegisterCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRegisterParams
}

type RequestMissionCall struct {
	Func   *wasmlib.ScFunc
	Params MutableRequestMissionParams
}

type SubmitLastMapCall struct {
	Func *wasmlib.ScFunc
}

type UploadMapCall struct {
	Func   *wasmlib.ScFunc
	Params MutableUploadMapParams
}

type GetMissionCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGetMissionResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) Init(ctx wasmlib.ScFuncCallContext) *InitCall {
	f := &InitCall{Func: wasmlib.NewScInitFunc(ctx, HScName, HFuncInit, keyMap[:], idxMap[:])}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) Register(ctx wasmlib.ScFuncCallContext) *RegisterCall {
	f := &RegisterCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRegister)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) RequestMission(ctx wasmlib.ScFuncCallContext) *RequestMissionCall {
	f := &RequestMissionCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncRequestMission)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) SubmitLastMap(ctx wasmlib.ScFuncCallContext) *SubmitLastMapCall {
	return &SubmitLastMapCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSubmitLastMap)}
}

func (sc Funcs) UploadMap(ctx wasmlib.ScFuncCallContext) *UploadMapCall {
	f := &UploadMapCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncUploadMap)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) GetMission(ctx wasmlib.ScViewCallContext) *GetMissionCall {
	f := &GetMissionCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetMission)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}
