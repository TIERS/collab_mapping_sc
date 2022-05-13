// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package controllerswarm

import "github.com/iotaledger/wasp/packages/vm/wasmlib/go/wasmlib"

type CompareMapCall struct {
	Func    *wasmlib.ScFunc
	Params  MutableCompareMapParams
	Results ImmutableCompareMapResults
}

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

type SubmitMapCall struct {
	Func   *wasmlib.ScFunc
	Params MutableSubmitMapParams
}

type GetMissionCall struct {
	Func    *wasmlib.ScView
	Params  MutableGetMissionParams
	Results ImmutableGetMissionResults
}

type GlobalMapCall struct {
	Func    *wasmlib.ScView
	Results ImmutableGlobalMapResults
}

type Funcs struct{}

var ScFuncs Funcs

func (sc Funcs) CompareMap(ctx wasmlib.ScFuncCallContext) *CompareMapCall {
	f := &CompareMapCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncCompareMap)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

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

func (sc Funcs) SubmitMap(ctx wasmlib.ScFuncCallContext) *SubmitMapCall {
	f := &SubmitMapCall{Func: wasmlib.NewScFunc(ctx, HScName, HFuncSubmitMap)}
	f.Func.SetPtrs(&f.Params.id, nil)
	return f
}

func (sc Funcs) GetMission(ctx wasmlib.ScViewCallContext) *GetMissionCall {
	f := &GetMissionCall{Func: wasmlib.NewScView(ctx, HScName, HViewGetMission)}
	f.Func.SetPtrs(&f.Params.id, &f.Results.id)
	return f
}

func (sc Funcs) GlobalMap(ctx wasmlib.ScViewCallContext) *GlobalMapCall {
	f := &GlobalMapCall{Func: wasmlib.NewScView(ctx, HScName, HViewGlobalMap)}
	f.Func.SetPtrs(nil, &f.Results.id)
	return f
}